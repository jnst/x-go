.PHONY: fmt lint test gen

fmt:
	go fmt ./...
	gofumpt -w .
	gci -w -local github.com/jnst/x-go .

lint:
	golangci-lint run

test:
	go test -v ./...

gen:
	protoc -I grpcapp/internal/proto \
		--go_out=grpcapp/internal/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=grpcapp/internal/gen \
		--go-grpc_opt=paths=source_relative \
		grpcapp/internal/proto/**/*.proto
