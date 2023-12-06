
all: solutions

.PHONY: solutions test

solutions:
	go run ./cmd/main.go

test:
	go test -race ./...

