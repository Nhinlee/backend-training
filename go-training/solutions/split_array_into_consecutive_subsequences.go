package solutions

import "math"

func isPossible(nums []int) bool {
	flags := make(map[int]int)
	var maxCount float64 = 0
	count := 0
	for _, num := range nums {
		flags[num]++
		maxCount = math.Max((float64(flags[num])), maxCount)
	}

	for _, v := range flags {
		if v == int(maxCount) {
			count++
		}
	}

	return count >= 3
}
