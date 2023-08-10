PKG_PREFIX := github.com/pixconf/pixconf

COMMIT_TAG ?= $(shell git rev-parse HEAD)

GO_BUILDINFO = -s -w \
	-X '$(PKG_PREFIX)/internal/buildinfo.Version=dev' \
	-X '$(PKG_PREFIX)/internal/buildinfo.Commit=$(COMMIT_TAG)'

.PHONY: $(MAKECMDGOALS)

help:
	@echo "read Makefile"

tests: test-lint test-unit

test-lint:
	golangci-lint run

test-unit:
	go test -race -cover ./...

test-pure:
	CGO_ENABLED=0 go test ./...

test-full:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

build-protos:
	rm -f internal/protos/*.pb.go
	protoc --go_out=internal/protos/ --go-grpc_out=internal/protos/ internal/protos/src/hub_service.proto

all: build-agent build-hub build-secrets

build-agent:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-agent-linux-amd64 $(PKG_PREFIX)/cmd/agent

build-hub:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-hub-linux-amd64 $(PKG_PREFIX)/cmd/hub

build-secrets:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-secrets-linux-amd64 $(PKG_PREFIX)/cmd/secrets

update:
	go get -v -u -d ./...
	go mod tidy -v -compat=1.20
