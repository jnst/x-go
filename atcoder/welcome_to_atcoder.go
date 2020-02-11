package main

import "fmt"

func main() {
	var a, b, c int
	var s string
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d %d", &b, &c)
	_, _ = fmt.Scanf("%s", &s)
	fmt.Printf("%d %s\n", a+b+c, s)
}
