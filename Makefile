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

test-unit-coverage:
	go test -coverprofile=coverage.out ./... -timeout 120s
	go tool cover -html=coverage.out

build-protos:
	rm -f internal/protos/*.pb.go
	protoc --go_out=internal/protos/ --go-grpc_out=internal/protos/ internal/protos/src/hub_service.proto
	protoc --go_out=internal/protos/ --go-grpc_out=internal/protos/ internal/protos/src/secrets.proto

all: build-agent build-hub build-secrets

build-agent:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-agent-linux-amd64 $(PKG_PREFIX)/cmd/agent

build-hub:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-hub-linux-amd64 $(PKG_PREFIX)/cmd/hub

build-secrets:
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GO_BUILDINFO)" -o build/pixconf-secrets-linux-amd64 $(PKG_PREFIX)/cmd/secrets

update:
	go get -v -u -d ./...
	go mod tidy -v -compat=1.20
