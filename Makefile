PKG_PREFIX := github.com/pixconf/pixconf

COMMIT_TAG ?= $(shell git rev-parse HEAD)

GO_BUILDINFO = -s -w \
	-X '$(PKG_PREFIX)/internal/buildinfo.Version=dev' \
	-X '$(PKG_PREFIX)/internal/buildinfo.Commit=$(COMMIT_TAG)'

.PHONY: $(MAKECMDGOALS)

help:
	@echo "read Makefile"

tests: test-lint test-race

test-lint:
	golangci-lint run

test-race:
	go test -race -cover ./...

test-pure:
	CGO_ENABLED=0 go test ./...

test-full:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

all: build-agent build-server

build: build-agent build-server

build-agent:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-agent-linux-amd64 $(PKG_PREFIX)/cmd/agent
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-agent-linux-arm64 $(PKG_PREFIX)/cmd/agent

build-server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-server-linux-amd64 $(PKG_PREFIX)/cmd/server
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-server-linux-arm64 $(PKG_PREFIX)/cmd/server

update:
	go get -v -u ./...
	go mod tidy -v -compat=1.23
