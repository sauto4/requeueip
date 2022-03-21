# IPAM CNI 插件调研（其一）

>来聊聊三个简单的 IPAM 项目：[static](https://www.cni.dev/plugins/current/ipam/static/)、[host-local](https://www.cni.dev/plugins/current/ipam/host-local/)（[CNI Plugin v1.1.0](https://github.com/containernetworking/plugins/releases/tag/v1.1.0)）和 [whereabouts](https://github.com/k8snetworkplumbingwg/whereabouts)（[v0.5.1](https://github.com/k8snetworkplumbingwg/whereabouts/releases/tag/v0.5.1)）。



## static

先说 static，它其实是一个非常非常简单的 IPAM 插件，代码量很小，给容器分配固定的 IPv4 和 IPv6 地址，它多用于容器的网络调试。简单理解就是，任何一个容器通过 static 来分配 IP，那么它们都会拿到同一个结果。

所以这个项目并不会体现很多我们比较关注的元素，比如锁、IP 分配算法、CRD 等；不过麻雀虽小，五脏俱全，它作为一个 CNI IPAM 插件的入门项目还是非常合格的。同时出于 static 主要由 CNI 容器网络工作小组维护的缘故，其源码总是会随着 CNI 协议（目前最新为 [1.0.1](https://github.com/containernetworking/cni/releases/tag/v1.0.1)）以及相关库的迭代而第一时间的更新， 使得其源码非常的 “时髦”（CNI Plugin 仓库中的其他插件同理）。

阅读 static 源码非常有助于理解以下内容：

- CNI 中约定的三类自定义参数（`CNI_ARGS`，`args conventions`，`capability argument`）的工作原理以及优先级。
- CNI 插件中关于 “执行结果” 字段 `prevResult` 的使用场景与方法。
- 关于 CNI 0.1.0/0.2.0 中约定的结果只能返回单一 IPv4 和 IPv6 地址的基本校验写法（CNI 0.3.0 之后都支持返回多 IP）。

其余就没有什么特别的亮点了，这里也不作展开说明了。



## host-local

### 机制

host-local 相较于 static 就会强大很多，它可以分配指定地址范围内的 IPv4 与 IPv6 地址给容器，然后将 IP 与容器 ID 的对应关系持久化到本地文件。值得注意的是，由于其直接将本地文件当数据库的做法，**使得它不能跨节点工作**。所以接下来我们对于 host-local 的讨论，都是仅限于单节点场景的。

先来看一个 host-local 的示例配置。

```json
{
    "ipam": {
        "type": "host-local",
        // Ranges
        "ranges": [
            // RangeSet
            [
                // Range
                {
                    "subnet": "10.10.0.0/16",
                    "rangeStart": "10.10.1.20",
                    "rangeEnd": "10.10.3.50",
                    "gateway": "10.10.0.254"
                },
                {
                    "subnet": "172.16.5.0/24"
                }
            ],
            [
                {
                    "subnet": "3ffe:ffff:0:01ff::/64",
                    "rangeStart": "3ffe:ffff:0:01ff::0010",
                    "rangeEnd": "3ffe:ffff:0:01ff::0020"
                }
            ]
        ],
        "dataDir": "/run/my-orchestrator/container-ipam-state"
    }
}
```

主要关注 host-local 所抽象出来的三个概念：Ranges，RangeSet，Range。它们呈一种包含关系，即一个 Ranges 中含有多个 RangeSet，一个 RangeSet 中含有多个 Range。

通俗的解释一下用法：

- `Range`：可以理解为一个 IP 范围的规则，多个 Range 可以搭配使用，多个 Range 所圈定的 IP 范围为他们的并集。需要注意的是，**Range 之间不允许有交集**。
- `RangeSet`：可以理解为一个 IP 池，池的范围由其持有的所有 Range 决定。IP 就是从一个个池中分出来的，如果定义了多个 IP 池，那么每次调用 IPAM 就能返回多个 IP。
- `Ranges`：为了适配 CNI 0.3.0 及以上版本中关于返回多 IP 的约定，废弃了原 `range` 字段（仍然支持，语法同上述示例 Json 中的 Range 块，层级需与 Ranges 保持一致），进而衍生出 `ranges` 字段。

所以在使用上述配置的情况下，当 host-local 被调用时会返回一个 IPv4 和一个 IPv6 地址，且 IPv4 地址的选择范围为两个 Range 的并集。这个模型就我个人看来是非常合理的，RequeueIP 的主要目标也是 Operator 化 host-local，让它可以**跨界点工作**的同时而且还无比**聪明**。



### 锁

#### 实现

host-local 采用**文件锁**来控制单节点上的 IP 分配并发问题。不同于一般的文件锁实现（比如利用 `mkdir` 同一个目录报错的机制），它使用 `flock` 系统调用原生的排他锁。`flock` 由 Linux 内核实现，它有三种操作类型：

- `LOCK_SH`：共享锁，多个进程可以使用同一把锁，常用于读锁。
- `LOCK_EX`：互斥锁，同时只允许一个进程使用，常用于写锁。
- `LOCK_UN`：释放锁。

最不错的地方是，`flock` 实现的文件锁的状态是与调用它的进程生命同步的，当异常发生而导致进程意外结束时，锁也会随之释放，无需考虑锁资源的相关回收问题。



#### 加锁方式

host-local 在操作 IP 的整个流程中，全局只会锁住一个文件（`{DATA_DIR}/{NETWORK_NAME}/lock`），所以可以理解为 IP 的分配与释放流程是全局串行的，即节点上同时间只会有一个 IP 在被操作。



### IP 分配算法

host-local 的具体 IP 分配流程如下（即 IP Allocator 的主要逻辑）：

1. 为每个 RangeSet 创建 Allocator，同时分配每个 Allocator 的 `rangeID` 编号（对应其在 Ranges 数组中的索引号）。
2. 遍历 Allocator，为每个 Allocator 初始化一个迭代器，用于在该 Allocator 对应的 RangeSet 中迭代 IP。
3. 初始化迭代器时，首先要通过 1 中传入的 `rangeID` 来从对应文件（`{DATA_DIR}/last_reserved_ip.{RANGE_ID}`）中检索起始游标，即开始迭代 IP 的起点。
4.  所以不难理解，每个 Allocator 都会有各自对应的 `rangeID` “游标文件”，**它用于存放该 Allocator 上一次成功分配的 IP**。如果在初始化迭代器时，检索到了对应 `rangeID` 的缓存游标文件，那就会从该文件中记录 IP 的下一个 IP 开始迭代；反之则从头开始迭代，即配置文件中 `rangeStart` 字段所对应的 IP（类似的思路，通过生成随机数来决定迭代的起点，一样能一定程度上的减少碰撞与并发）。
5. 每个 Allocator 在成功分出一个 IP 后，进而尝试在数据目录落盘，流程结束。

host-local 中，IP 的分配关系是通过在数据目录中存 **IP  地址名称的文件**来维护的，IP 文件中的内容为：`{CNI_CONTAINERID} + "\r\n" + {CNI_IFNAME}`。（`TODO`：为什么要刻意去存其实不太相关的设备名？猜测可能与 multus 相关）



### 其他

- host-local 关于 DNS 配置功能的部分，提供了基于 resolv.conf 配置文件解析的实现，可通过 `resolvConf` 字段声明定义。
- host-local 是一定存在 IP 泄漏现象的（一个例子 [issue](https://github.com/containernetworking/plugins/issues/498)），从代码的层面来看也很好理解，在创建完 IP 文件后，且还未写入容器 ID 与设备信息建立绑定关系时，某些故障发生了，这时就会出现内容为空的 IP 文件，这是代码层面难以避免的。一定程度上来说，每个 IPAM 插件都应该要有 IP GC 的逻辑或相关组件（声明式或许不需要？）。



### 小结

1. 网络相关的轮子丰富强悍，关于 Range、RangeSet 的入参校验逻辑非常值得借鉴。
2. 基于系统调用 `flock` 的文件锁实现值得借鉴。
3. 采用迭代器设计模式实现的 IP Allocator 值得借鉴。
4. 存储后端与 IP 分配器解耦的思想，应该严格遵循。



## whereabouts

### 机制

whereabouts 相较于 host-local 来说，最大的却别就在于它支持跨节点的 IP 分配。思路也非常简单，将存储数据的后端替换为一个集群维度而非单节点的实现即可，whereabouts 提供了两种存储后端实现：etcd 和 K8s CRD。

whereabouts 与 RequeueIP 的最大不同在于，whereabouts 仅仅把 CRD 当成了一个纯粹的 “数据库”，它没有 Operator 的模型与思想。

先来看一个基本配置示例：

```json
{
    "ipam": {
        "type": "whereabouts",
        "range": "192.168.2.225-192.168.2.250/28",
        "exclude": [
            "192.168.2.229/30",
            "192.168.2.236/32"
        ]
    }
}
```

1. 非常直观的可以看到，whereabouts 并没有抽象出 CNI 0.3.0 之后约定的 *”ranges“* 概念，还是沿用 `range` 字段，这说明它一次只能分出一个 IPv4 或 IPv6 的地址，这是比较落后的。然后就是增强了一个 `exclude` 语法，用于更灵活的定义 IP 池的范围。（host-local 一样可以通过定义多个 Range 块的形式来描述这个语义）

2. 相较于 host-local，whereabouts 的模型就只抽象出了 IP 池的概念，用户可以通过诸如 `range`、`exclude`、`range_start`、`range_end` 字段来灵活的定义这个 IP 池的范围，进而从该范围内分配与管理 IP。同时，一旦该节点通过 Network Configuration 的 `range` 字段指定了 IP 池的范围，那被调度至该节点的 *Pod* 便只能从该池中分配 IP，whereabouts 是没有什么选池逻辑的，配置怎么配，当前节点就怎么用指定的 IP 池。

3. 当然，可以通过在不同的节点上定义不同的 Network Configuration 来达到集群中有多个 IP 池的状态，理论上这些 IP 池是**不能有交集**的。whereabouts 的 CRD 后端实现单独针对上述存在交集的场景进行了相关增强，而 etcd 后端却不支持该能力。
4. whereabouts 的代码开始于 static，其设计参考了早期的 host-local。



### CRD

whereabouts 中最核心的 CRD 就是 *IPPool* 。

```go
type IPPoolSpec struct {
    Range string `json:"range"`                                 // CIDR
    Allocations map[string]IPAllocation `json:"allocations"`    // IP 分配关系
}

type IPAllocation struct {
    ContainerID string `json:"id"`
    PodRef      string `json:"podref,omitempty"`                // "namespace/podname" 字符串
}
```

其字段言简意赅，`Range` 约定了该 IP 池的范围，`Allocations` 记录了 IP 分配的情况。在常规的 IP 分配逻辑中，大体的流程即：

1. 根据入参从 K8s 中查询相关 *IPPool* 资源。
2. 分配 IP，记录分配情况。
3. 将修改后的 *IPPool* 写回 K8s。

非常的简单明了。



第二个 CRD 是 *OverlappingRangeIPReservation*，它就用于处理之前我们上面所提到的多个 IP 池存在交集的场景。

```go
type OverlappingRangeIPReservationSpec struct {
    ContainerID string `json:"containerid"`         // 容器 ID
    PodRef      string `json:"podref,omitempty"`    // "namespace/podname" 字符串
}
```

光看 *OverlappingRangeIPReservation* 的 `Spec` 字段你肯定理解不了这个 CRD 有什么用，不过当你发现 *OverlappingRangeIPReservation* 通常以一个 IP 地址命名时，就能立马明白它是用于记录一个 IP 具体分配情况的资源，而且是在集群范围内的记录。

在常规的 IP 分配场景中，各节点都基于各自所定义的 *IPPool* 来分配与管理 IP，它们之间是**无感知**的。当两个节点分别定义了不同的 *IPPool*，且两者所圈定的范围相交时，那么就有可能出现位于同一集群中两个不同节点的 *Pod* 持有相同 IP 的场景，这肯定是会有问题的。

而当通过 whereabouts 配置中的 `enable_overlapping_ranges` 参数打开了该特性后，每次 IP 分配的结果就会以 *OverlappingRangeIPReservation* 资源的形式来持久化。再回到上面那个发生 IP 冲突的场景，当其中一个 IP 已经被分给了某个 *Pod* 后，另一个节点上的 *Pod* 在分配 IP 时则会检索 *OverlappingRangeIPReservation* 资源，进而知晓该 IP 已被占用。



### 锁

在 whereabouts 的 CRD 后端实现中，通过以下机制处理并发：

1. 初始化 IPAM CRD 后端时，通过 *Lease* 资源锁进行选主。即集群内同一时间内仅有一个 *Pod* 在分配 IP。
2. JSON Patch *IPPool* 资源时，会比较 `resourceVersion` 字段。



### IP 分配算法

简单的 `for` 循环，由对应 IP 池的 `range_start` 遍历至 `range_end`。分配出 IP 后则交由对应的后端进行存储。



### IP GC

TODO



### 小结

1. whereabouts 的代码比较凌乱。
2. whereabouts 存在一些 CRD 的实践，却没有 Operator 的任何逻辑。CRD 之于 whereabouts 也就是换了个地方存数据。





