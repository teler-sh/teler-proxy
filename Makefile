APP_NAME = teler-proxy
VERSION  = $(shell git describe --always --tags)

GO_LDFLAGS = "-s -w -X 'github.com/kitabisa/teler-proxy/common.Version=${VERSION}'"

vet:
	go vet ./...

lint:
	golangci-lint run ./...

semgrep:
	semgrep --config auto

bench:
	go test ./pkg/tunnel/... -bench . -cpu 4 -benchmem $(ARGS)

cover: FILE := /tmp/teler-coverage.out # Define coverage file
cover: ## Runs the tests and check & view the test coverage
	go test -race -coverprofile=$(FILE) -covermode=atomic ./...
	go tool cover -func=$(FILE)

pprof: ARGS := -cpuprofile=cpu.out -memprofile=mem.out -benchtime 30s
pprof: bench
pprof:
	cp cpu.out default.pgo

test:
	go test -race -v ./pkg/tunnel/...

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

all: test report build