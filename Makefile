.PHONY: fmt test

fmt:
	go fmt ./...
	goimports -w -local github.com/jnst/x-go .

test:
	go test -v ./...
