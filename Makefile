.PHONY: fmt test

fmt:
	go fmt ./...
	goimports -w -local github.com/jnst/x-go .

lint:
	golangci-lint run

test:
	go test -v ./...
