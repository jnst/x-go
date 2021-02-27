package leetcode

// Merge merges nums1 and nums2 as one sorted array.
// The length of nums1 is m+n by constraint.
func Merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
