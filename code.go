package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Pair represents key value pair.
type Pair struct {
	K string
	V int
}

// input:
// [
//   ("news.example.com", 100),
//   ("example.com", 220),
//   ("google.com", 500),
//   ("example.jp", 30)
// ]
// output:
// {
//   "com": 820,
//   "example.com": 320,
//   "news.example.com": 100,
//   "google.com": 500,
//   "jp": 30,
//   "example.jp": 30,
// }
func SumEachLevelDomain(pairs []Pair) map[string]int {
	m := map[string]int{}

	for _, pair := range pairs {
		levelDomains := strings.Split(pair.K, ".")
		size := len(levelDomains)
		for i := range levelDomains {
			domain := levelDomains[size - i - 1:]
			target := strings.Join(domain, ".")
			count, ok := m[target]
			if ok {
				m[target] = count + pair.V
			} else {
				m[target] = pair.V
			}

			fmt.Printf("%v => %s: %d\n", levelDomains, target, m[target])
		}
	}

	return m
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	arr := strings.Split(s, "")

	size := len(arr)
	for i := range arr {
		if arr[i] != arr[size-1-i] {
			return false
		}
	}

	return true
}

func ThreeSum(nums []int) [][]int {
	var result [][]int

	for i, n1 := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			for k, n3 := range nums {
				if i == k || j == k {
					continue
				}
				target := nums[i] + nums[j] + nums[k]
				if target == 0 {
					arr := []int{n1, n2, n3}
					result = append(result, arr)
				}
			}
		}
	}

	return result
}

// 階乗を求めるコード
func factorial(n uint64) uint64 {
	result := uint64(1)
	for i := uint64(0); i < n; i++ {
		result *= n - i
	}
	return result
}

func main() {
	for i := 0; i < 30; i++ {
		fmt.Printf("%2d\t%d\n", i, factorial(uint64(i)))
	}
}
