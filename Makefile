all: lint

run:
	go run ./cmd/main.go
lint: 
	golangci-lint run --fix 
