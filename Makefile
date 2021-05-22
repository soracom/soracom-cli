VERSION ?= 0.0.0
GOOS ?= darwin
GOARCH ?= amd64
EXT ?= 

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

install: ## Install dependencies
	@echo 'Installing build dependencies ...'
	go get -u golang.org/x/tools/cmd/goimports
	@echo 'Installing commands used with "go generate" ...'
	go get -u github.com/elazarl/goproxy
	go mod tidy
.PHONY:install

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
.PHONY:generate

build: ## Build codes
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build \
		-ldflags="-X 'github.com/soracom/soracom-cli/soracom/generated/cmd.version=$(VERSION)'" \
		-o "soracom/dist/$(VERSION)/soracom_$(VERSION)_$(GOOS)_$(GOARCH)$(EXT)" \
		./soracom
.PHONY:build
