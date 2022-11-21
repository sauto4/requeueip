# Copyright 2022 The RequeueIP Authors.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# ==============================================================================
# Makefile helper functions for generate necessary files
#

.PHONY: gen.run
#gen.run: gen.manifests gen.generate-k8s-api
gen.run: gen.clean gen.manifests gen.generate-k8s-api

.PHONY: gen.manifests
gen.manifests:
	@echo "===========> Generate ClusterRole and CustomResourceDefinition objects."
	(cd ${ROOT_DIR} && ./tools/k8s-api-gen/controller-gen.sh manifests)

.PHONY: gen.generate-k8s-api
gen.generate-k8s-api:
	@echo "===========> Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations."
	(cd ${ROOT_DIR} && ./tools/k8s-api-gen/controller-gen.sh deepcopy)

.PHONY: gen.clean
gen.clean:
	@rm -rf @${ROOT_DIR}/api
	@rm -rf @${ROOT_DIR}/charts
	@$(FIND) -type f -name '*_generated.go' -delete