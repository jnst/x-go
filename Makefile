.PHONY: fmt lint test

fmt:
	go fmt ./...
	gofumpt -w .
	gci -w -local github.com/jnst/x-go .

lint:
	golangci-lint run

test:
	go test -v ./...
