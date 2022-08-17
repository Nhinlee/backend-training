package solutions

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, index := 0, 0, 0
	temp := make([]int, m+n+1)
	for i < m && j < n {
		if nums1[i] > nums2[j] {
			temp[index] = nums2[j]
			j++
		} else {
			temp[index] = nums1[i]
			i++
		}
		index++
	}
	for ; i < m; i++ {
		temp[index] = nums1[i]
		index++
	}
	for ; j < n; j++ {
		temp[index] = nums2[j]
		index++
	}
	copy(nums1, temp)
}
