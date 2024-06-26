PROJECT_NAME := "kube-admin"
MAIN_FILE := "main.go"
OUTPUT_NAME=${PROJECT_NAME}

## 需要跟go.mod保持一致
PKG := "github.com/xiaoxin1992/kube-admin"


PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)


BUILD_VERSION := $(shell git describe --tags `git rev-list --tags --max-count=1`)
BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "${PKG}/version"


.PHONY: all dep lint vet test-converage build clean


all: build

dep:	## Get the dependencies
	@go mod tidy

lint:	## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet:	## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

build: dep ## Build the binary file
	@go build -a -o dist/${OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitTag=${BUILD_VERSION}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'" ${MAIN_FILE}

linux: dep ## Build the binary file
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GitBranch=${BUILD_BRANCH}' -X '${VERSION_PATH}.GitTag=${BUILD_VERSION}' -X '${VERSION_PATH}.GitCommit=${BUILD_COMMIT}' -X '${VERSION_PATH}.BuildTime=${BUILD_TIME}' -X '${VERSION_PATH}.GoVersion=${BUILD_GO_VERSION}'" ${MAIN_FILE}

run: dep build ## Run Server
	@./dist/${PROJECT_NAME}

clean: ## Remove previous build
	@go clean .
	@rm -f dist/${PROJECT_NAME}


#gen: ## Init Service
#	@protoc -I=.  -I=/usr/local/include --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} apps/*/pb/*.proto
#	@protoc-go-inject-tag -input=apps/*/*.pb.go

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'