.PHONY: format

format:
	go fmt ./...
	goimports -w -local github.com/jnst/x-go .
