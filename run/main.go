package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	s = append(s, 1)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
