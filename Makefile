CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin


all: lint

run: check-go
	@go run ./cmd/main.go

lint: check-go check-lint
	@golangci-lint run --fix 

build: check-go bindir
	@go build -o ${BINDIR}/app ./cmd/main.go

bindir:
	@mkdir -p ${BINDIR}

check-go:
	@command -v go >/dev/null 2>&1 || { echo >&2 "Go is not installed. Please install Go to proceed."; exit 1; }

check-lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo >&2 "golangci-lint is not installed. Please install golangci-lint to proceed."; exit 1; }
