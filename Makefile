APP_NAME = teler-proxy
VERSION  = $(shell git describe --always --tags)

GO_MOD_VERSION := $(shell grep -Po '^go \K([0-9]+\.[0-9]+(\.[0-9]+)?)$$' go.mod)
GO := go${GO_MOD_VERSION}
GO_LDFLAGS = "-s -w -X 'github.com/teler-sh/teler-proxy/common.Version=${VERSION}'"

ifeq ($(shell which ${GO}),)
	GO = go
endif

vet:
	$(GO) vet ./...

lint:
	golangci-lint run --tests=false ./...

semgrep:
	semgrep --config auto

bench:
	$(GO) test ./pkg/tunnel/... -run "^$$" -bench . -cpu 4 -benchmem $(ARGS)

cover: FILE := /tmp/teler-coverage.out # Define coverage file
cover: PKG := ./pkg/tunnel/...
cover:
	$(GO) test -race -coverprofile=$(FILE) -covermode=atomic $(PKG)
	$(GO) tool cover -func=$(FILE)

pprof: ARGS := -cpuprofile=cpu.out -memprofile=mem.out -benchtime 30s
pprof: bench

pgo: pprof
pgo:
	cp cpu.out default.pgo

test:
	$(GO) test -race -v ./pkg/tunnel/...

test-all: test vet lint semgrep

report:
	goreportcard-cli

build:
	@echo "Building binary"
	@mkdir -p bin/
	CGO_ENABLED="1" go build -ldflags ${GO_LDFLAGS} -trimpath $(ARGS) -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

build-pgo: ARGS := -pgo=$(shell pwd)/default.pgo
build-pgo: build

docker:
	@echo "Building image"
	docker build -t ${APP_NAME}:latest --build-arg="VERSION=${VERSION}" .

clean:
	@echo "Removing binaries"
	@rm -rf bin/

teler-proxy: build

ci: vet build clean

all: test report build