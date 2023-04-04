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
gen.run: gen.clean gen.manifests gen.deepcopy gen.openapi

.PHONY: gen.clean
gen.clean: gen.manifests-clean gen.openapi-clean

.PHONY: gen.manifests
gen.manifests:
	@echo "===========> Generate ClusterRole and CustomResourceDefinition"
	@(cd ${ROOT_DIR} && ./tools/scripts/controller-gen.sh manifests)

.PHONY: gen.manifests-verify
gen.manifests-verify:
	@echo "===========> Verify ClusterRole and CustomResourceDefinition"
	@(cd ${ROOT_DIR} && ./tools/scripts/controller-gen.sh verify)

.PHONY: gen.manifests-clean
gen.manifests-clean:
	@echo "===========> Clean ClusterRole and CustomResourceDefinition"
	@(cd ${ROOT_DIR} && ./tools/scripts/controller-gen.sh clean)
	
.PHONY: gen.deepcopy
gen.deepcopy:
	@echo "===========> Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations"
	@(cd ${ROOT_DIR} && ./tools/scripts/controller-gen.sh deepcopy)

.PHONY: gen.openapi-validate
gen.openapi-validate:
	@echo "===========> Validate OpenAPI spec openapi.yaml"
	@(cd ${ROOT_DIR} && ./tools/scripts/swagger-gen.sh validate)

.PHONY: gen.openapi
gen.openapi: gen.openapi-validate
	@echo "===========> Generate OpenAPI source code"
	@(cd ${ROOT_DIR} && ./tools/scripts/swagger-gen.sh gen)

.PHONY: gen.openapi-verify
gen.openapi-verify: gen.openapi-validate
	@echo "===========> Verify OpenAPI source code"
	@(cd ${ROOT_DIR} && ./tools/scripts/swagger-gen.sh verify)

.PHONY: gen.openapi-clean
gen.openapi-clean:
	@echo "===========> Clean OpenAPI source code"
	@(cd ${ROOT_DIR} && ./tools/scripts/swagger-gen.sh clean)
