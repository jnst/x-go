package code

import (
	"strconv"
	"strings"
)

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
