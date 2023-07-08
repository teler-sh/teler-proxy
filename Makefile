APP_NAME = teler-proxy
VERSION  = $(shell git describe --always --tags)

vet:
	go vet ./...

lint:
	golangci-lint run ./...

semgrep:
	semgrep --config auto

test: vet lint semgrep

report:
	goreportcard-cli

build:
	@echo "Building binary"
	@mkdir -p bin/
	CGO_ENABLED="1" go build -ldflags "-s -w -X github.com/kitabisa/teler-proxy/common.Version=${VERSION}" -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

clean:
	@echo "Removing binaries"
	@rm -rf bin/

teler-proxy: build

all: test report build