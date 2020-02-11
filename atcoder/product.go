package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	if (a * b) % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
