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

DEP_TOOLS ?= gsemver golines git-chglog github-release go-mod-outdated golangci-lint goimports
OTHER_TOOLS ?= depth go-callvis gothanks

tools.install: $(addprefix tools.install., $(DEP_TOOLS), ${OTHER_TOOLS})
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

.PHONY: install.gsemver
install.gsemver:
	@$(GO) install github.com/arnaud-deprez/gsemver@latest

.PHONY: install.git-chglog
install.git-chglog:
	@$(GO) install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

.PHONY: install.github-release
install.github-release:
	@$(GO) install github.com/github-release/github-release@latest

.PHONY: install.golines
install.golines:
	@$(GO) install github.com/segmentio/golines@latest

.PHONY: install.go-mod-outdated
install.go-mod-outdated:
	@$(GO) install github.com/psampaz/go-mod-outdated@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.depth
install.depth:
	@$(GO) install github.com/KyleBanks/depth/cmd/depth@latest

.PHONY: install.go-callvis
install.go-callvis:
	@$(GO) install github.com/ofabry/go-callvis@latest

.PHONY: install.gothanks
install.gothanks:
	@$(GO) install github.com/psampaz/gothanks@latest