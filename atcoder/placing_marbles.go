package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(strings.Count(s, "1"))
}
