package main

import (
	"flag"
	"github.com/jnst/x-go/sample"
	"github.com/jnst/x-go/std"
	"strings"

	"github.com/jnst/x-go/code"
)

func main() {
	//f := flag.Int64("n", 0, "factorial -n 10")
	s := flag.String("e", "", "execute any code")

	switch strings.ToLower(*s) {
	case "factorial":
		code.Factorial(10)
	case "slack":
		sample.Post()
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
