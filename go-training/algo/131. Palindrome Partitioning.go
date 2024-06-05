package algo

func partition(str string) [][]string {
	result := [][]string{}
	path := []string{}
	backtrack(str, 0, path, &result)
	return result
}

func backtrack(str string, start int, path []string, result *[][]string) {
	if start == len(str) {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	for end := start + 1; end <= len(str); end++ {
		if isPalindrome(str, start, end-1) {
			path = append(path, str[start:end])
			backtrack(str, end, path, result)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
