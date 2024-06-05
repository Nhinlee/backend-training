package algo

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	rs := 0

	for left < right {
		// Calculate the area
		width := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := width * h
		if area > rs {
			rs = area
		}
	}

	return rs
}
