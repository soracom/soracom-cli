VERSION ?= 0.0.0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
ifeq ($(GOOS),windows)
	EXT = .exe
else
	EXT =
endif
OUTPUT ?= soracom/dist/$(VERSION)/soracom_$(VERSION)_$(GOOS)_$(GOARCH)$(EXT)

GOCYCLO_OVER ?= 23

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

install-dev-deps: ## Install dev dependencies
	@echo 'Installing dependencies for development'
	go get -u \
		golang.org/x/lint/golint \
		github.com/fzipp/gocyclo/cmd/gocyclo
.PHONY:install-dev-deps

install-deps: ## Install dependencies
	@echo 'Installing build dependencies ...'
	go get -u golang.org/x/tools/cmd/goimports
	@echo 'Installing commands used with "go generate" ...'
	go get -u github.com/elazarl/goproxy
	go mod tidy
.PHONY:install-deps

test: ## Test generator's library
	@echo "Testing generator's source ..."
	go test ./generators/cmd/src
	go test ./generators/lib
.PHONY:test

generate: ## Generate source code for soracom-cli
	echo 'Generating generator ...'
	cd ./generators/cmd/src && \
	go generate && \
	go vet && \
	goimports -w ./*.go && \
	go build -o generate-cmd
	@echo 'Generating source codes for soracom-cli by using the generator ...'
	./generators/cmd/src/generate-cmd -a generators/assets/soracom-api.en.yaml -s generators/assets/sandbox/soracom-sandbox-api.en.yaml -t generators/cmd/templates -p generators/cmd/predefined -o soracom/generated/cmd/
	@echo 'Copying assets to embed ...'
	cp -r generators/assets/ soracom/generated/cmd/assets/
	make format
.PHONY:generate

build: ## Build codes
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-ldflags="-X 'github.com/soracom/soracom-cli/soracom/generated/cmd.version=$(VERSION)'" \
		-o $(OUTPUT) \
		./soracom
.PHONY:build

cross-build:
	for os in $(OS_LIST); do \
		for arch in $(ARCH_LIST); do \
			make build GOOS=$$os GOARCH=$$arch OUTPUT=soracom/dist/ghactions/$$os-$$arch/soracom$(EXT); \
		done; \
	done
.PHONY:cross-build

ci-build-artifacts: install-dev-deps install-deps generate metrics-gocyclo test lint ## Run `build-artifacts` action
	make cross-build OS_LIST="linux" ARCH_LIST="amd64 arm64 386 arm"
	make cross-build OS_LIST="darwin" ARCH_LIST="amd64 arm64"
	make cross-build OS_LIST="windows" ARCH_LIST="amd64 386" EXT=".exe"
	make cross-build OS_LIST="freebsd" ARCH_LIST="amd64 386"
.PHONY:ci-build-artifacts

format: ## Format codes
	go fmt ./...
.PHONY:format

lint: ## Lint codes
	golint -set_exit_status ./...
.PHONY:lint

metrics-gocyclo: ## Metrics with gocyclo
	gocyclo -over $(GOCYCLO_OVER) .
.PHONY:metrics-gocyclo
