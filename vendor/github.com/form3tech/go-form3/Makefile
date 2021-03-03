.DEFAULT_GOAL := default

swagger_codegen_version := "0.15.0"

platform := $(shell uname)
ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: test

build: goimports
	go install

install: install-aws-cli install-goimports install-swagger

install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

install-aws-cli:
	@echo "Installing aws-cli..."
	@sudo pip install awscli

install-goimports:
	@echo Installing goimports...
	@go get golang.org/x/tools/cmd/goimports

generate-swagger-model:
	@echo "Removing previously generated client"
	@rm -rf ./client ./models
	@echo "Generating new client"
	@swagger generate client -f ./swagger.yaml

GOFMT_FILES?=$$(find ./ -name '*.go' | grep -v vendor)

goimports:
	goimports -w $(GOFMT_FILES)

test: build
	go test $(TEST) -v $(TESTARGS) -timeout 120m

.PHONY: build test vet fmt fmtcheck errcheck vendor-status
