.PHONY: server

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME="bin/tidb-tools-ops"
MAIN_FILE := $(find . -name "main.go")
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: cmd

default: build
all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_FILE)
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