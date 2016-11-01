package main

import (
	"github.com/jnst/go-world/io"
	"github.com/jnst/go-world/protobuf"
	"github.com/jnst/go-world/redis"
)

func main() {
	io.PrintFileByIOUtil("io/sample.txt")
	protobuf.Check()
	redis.Connect()
}
