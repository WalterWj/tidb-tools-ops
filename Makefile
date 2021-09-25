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
REPO := github.com/WalterWj/tidb-tools-ops
_build := $(shell git rev-parse HEAD)
_GOVERSION := $(shell go version)
_BUILDTIME := $(shell date -d today +'%Y-%m-%d %T')
_VERSION := $(shell git symbolic-ref --short HEAD)
LDFLAGS := 
LDFLAGS += "-X $(REPO)/common/version.build=$(_build)
LDFLAGS += -X '$(REPO)/common/version.goVersion=$(_GOVERSION)'
LDFLAGS += -X '$(REPO)/common/version.version=$(_VERSION)'
LDFLAGS += -X '$(REPO)/common/version.buildTime=$(_BUILDTIME)'" 
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
	$(GOGET) github.com/Walterwj/tidb-tools-ops

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v