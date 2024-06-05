package algo

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Check if mid element is greater than the rightmost element
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
