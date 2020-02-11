package main

import (
	"fmt"
)

func main() {
	var size, n int
	_, _ = fmt.Scan(&size)

	var ns []int
	for i := 0; i < size; i++ {
		_, _ = fmt.Scan(&n)
		ns = append(ns, n)
	}

	_, count := loop(ns, 0)
	fmt.Println(count)
}

func loop(ns []int, count int) ([]int, int) {
	if ns == nil {
		return nil, count
	}

	arr := make([]int, 0, len(ns))
	for _, n := range ns {
		if isOdd(n) {
			return nil, count
		}
		arr = append(arr, n/2)
	}

	return loop(arr, count+1)
}

func isOdd(n int) bool {
	if n % 2 != 0 {
		return true
	}
	return false
}
