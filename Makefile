DEFAULT: build

MAKEFLAGS = --silent

# ============= BUILDS =============
.PHONY: build
build: ## build executable for current environment
	@./scripts/build

.PHONY: build-all
build-all: ## build all executables
	@./scripts/build linux
	@./scripts/build windows
	@./scripts/build macos

.PHONY: build-linux
build-linux: ## build executable for Linux
	@./scripts/build linux

.PHONY: build-windows
build-windows: ## build executable for Windows
	@./scripts/build windows

.PHONY: build-macos
build-macos: ## build executable for MacOs
	@./scripts/build macos

# ============= TOOLING =============
.PHONY: clean
clean: ## remove build artifacts
	rm -rf ./build/*
	@go clean ./...

.PHONY: install
install: ## installing tools / dependencies
	@go mod install

.PHONY: help
help: ## print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)