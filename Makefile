export CGO_ENABLED=0
export GO111MODULE=on

PARENT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)

.PHONY: build

all: build

build: # @HELP build the source code
build:
	go build -v ./...

test: # @HELP run the unit tests and source code validation
test: build license_check linters
	go test github.com/atomix/atomix-go-client/pkg/...

coverage: # @HELP generate unit test coverage data
coverage: build linters license_check
	./build/bin/coveralls-coverage

primitives: # @HELP compile the protobuf files (using protoc-go Docker)
	docker run -it \
		-v $(PARENT_DIR)/atomix-api:/go/src/github.com/atomix/atomix-api \
		-v `pwd`:/go/src/github.com/atomix/atomix-go-client \
		-w /go/src/github.com/atomix/atomix-go-client \
		--entrypoint build/bin/generate-primitives.sh \
		atomix/protoc-gen-atomix:latest

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

license_check: # @HELP examine and ensure license headers exist
	./build/bin/license-check
