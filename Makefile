.PHONY: server

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME="bin/tidb-tools-ops"
MAIN_FILE := $(shell find . -name "main.go")
BINARY_UNIX=$(BINARY_NAME)_unix

# build info
REPO := tidb-tools-ops
_build := $(shell git rev-parse HEAD)
_GOVERSION := $(shell go version)
_BUILDTIME := $(shell date -d today +'%Y-%m-%d %T')
_VERSION := $(shell git symbolic-ref --short HEAD)
# 编译中传入：版本，go 版本，git hash，build 时间
LDFLAGS := 
LDFLAGS += "-X $(REPO)/internal/version.build=$(_build)
LDFLAGS += -X '$(REPO)/internal/version.goVersion=$(_GOVERSION)'
LDFLAGS += -X '$(REPO)/internal/version.version=$(_VERSION)'
LDFLAGS += -X '$(REPO)/internal/version.buildTime=$(_BUILDTIME)'" 
# flags="-X main.cmd.build=`git rev-parse HEAD` -X main.cmd.goVersion=`go version` -X main.cmd.buildTime=`date -d today +'%Y-%m-%d %T'`"

.PHONY: cmd

default: build
all: test build
build:
	$(GOBUILD) -ldflags $(LDFLAGS) -o $(BINARY_NAME) -v $(MAIN_FILE)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GORUN) -n $(MAIN_FILE)
deps:
	$(GOGET) tidb-tools-ops

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v