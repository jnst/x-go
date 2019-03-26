.PHONY: format

format:
	gofmt -w *.go
	goimports -w *.go
