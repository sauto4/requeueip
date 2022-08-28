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

.PHONY: dependencies.run
dependencies.run: dependencies.critical dependencies.prefer

.PHONY: dependencies.critical
dependencies.critical: go.build.verify go.lint.verify release.gsemver.verify

.PHONY: dependencies.prefer
dependencies.prefer: release.git-chglog.verify release.github-release.verify