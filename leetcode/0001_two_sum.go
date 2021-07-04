package leetcode

// TwoSum returns indices of the nums array when the sum of the two numbers is
// equal to target.
func TwoSum(nums []int, target int) []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
