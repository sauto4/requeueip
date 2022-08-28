#!/usr/bin/make -f

.PHONY: manifests
manifests:
	@echo "Generate ClusterRole and CustomResourceDefinition objects."
	@tools/k8s-api-gen/controller-gen.sh manifests

.PHONY: generate-k8s-api
generate-k8s-api:
	@echo "Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations."
	@tools/k8s-api-gen/controller-gen.sh deepcopy