package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jnst/x-go/sample"
	"github.com/jnst/x-go/solana"
	"github.com/jnst/x-go/std"

	"github.com/jnst/x-go/code"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "specify the command name as argument")
		return
	}

	switch command := strings.ToLower(args[0]); command {
	case "factorial":
		code.Factorial(10)
	case "slack":
		sample.Post()
	case "solana-keygen":
		solana.GenerateVanityAddress(args[1])
	case "sort":
		std.PrintSortedHeroes()
		std.PrintSortedYearMonths()
	case "uid":
		sample.PrintXID()
		sample.PrintULID()
	default:
		flag.Usage()
	}
}
