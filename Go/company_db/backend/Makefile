# Go parameters
GOFMT := go fmt ./...
GOIMPORTS := goimports -w .
GOVET := go vet ./...
LINT := golangci-lint run
STATICCHECK := staticcheck ./...
MODTIDY := go mod tidy
BUILD := go build -o bin/app
CONCURRENCY := $(shell sysctl -n hw.ncpu)

.PHONY: all format vet lint staticcheck tidy build test clean build-linux build-darwin build-windows

all: format vet lint staticcheck tidy build

format:
	@echo "==> Formatting code with go fmt"
	$(GOFMT) || exit 1
	@echo "==> Cleaning imports with goimports"
	$(GOIMPORTS) || exit 1

vet:
	@echo "==> Running go vet"
	go vet ./... || exit 1

lint:
	@echo "==> Running golangci-lint"
	golangci-lint run --concurrency=$(CONCURRENCY) || exit 1

staticcheck:
	@echo "==> Running staticcheck"
	$(STATICCHECK) || exit 1

tidy:
	@echo "==> Tidying up go.mod and go.sum"
	$(MODTIDY) || exit 1

build:
	@echo "==> Building the application"
	mkdir -p bin
	$(BUILD) || exit 1

test:
	@echo "==> Running tests"
	go test ./... || exit 1

clean:
	@echo "==> Cleaning build artifacts"
	rm -rf bin

build-linux:
	GOOS=linux GOARCH=amd64 $(BUILD)

build-darwin:
	GOOS=darwin GOARCH=amd64 $(BUILD)

build-windows:
	GOOS=windows GOARCH=amd64 $(BUILD).exe
