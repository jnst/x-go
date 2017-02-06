package main

import (
	"flag"
	"fmt"
	"os"
)

// Factorial returns factorial number
func Factorial(n int64) int64 {
	if n <= 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	f := flag.Int64("n", 0, "factorial -n 10")
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.Usage()
		return
	}

	fmt.Printf("%d -> %d", *f, Factorial(*f))
}
