package main

import (
	"fmt"
	"strings"
)

func scan() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(strings.Count(s, "1"))
}

func scanf() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if (a*b)%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
