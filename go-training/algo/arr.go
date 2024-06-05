package algo

func specialArray(nums []int) int {
	// C1
	// sort.Ints(nums)
	// n := len(nums)

	// for i := 0; i <= n; i++ {
	// 	if i == 0 {
	// 		if nums[0] >= n {
	// 			return n
	// 		}
	// 	} else {
	// 		if nums[i] >= n-i && nums[i-1] < n-i {
	// 			return n - i
	// 		}
	// 	}
	// }

	// return -1

	// C2
	n := len(nums)
	freq := make([]int, n+1)

	for _, e := range nums {
		freq[min(e, n)]++
	}

	count := 0
	for i := n; i >= 0; i-- {
		count += freq[i]
		if i == count {
			return i
		}
	}

	return -1
}

func scoreOfString(s string) int {
	score := 0
	for i := 0; i < len(s)-1; i++ {
		tempScore := int(s[i]) - int(s[i+1])
		if tempScore < 0 {
			tempScore = -tempScore
		}
		score += tempScore
	}
	return score
}

// https://leetcode.com/problems/find-common-characters/?envType=daily-question&envId=2024-06-05
func commonChars(words []string) []string {
	var globalChars [26]int

	for _, ch := range words[0] {
		globalChars[int8(ch)-int8('a')]++
	}

	for _, word := range words {
		var chars [26]int
		for _, ch := range word {
			chars[int8(ch)-int8('a')]++
		}

		for i := 0; i < 26; i++ {
			globalChars[i] = min(globalChars[i], chars[i])
		}
	}

	rs := make([]string, 0)
	for i := 0; i < 26; i++ {
		if globalChars[i] > 0 {
			for j := 0; j < globalChars[i]; j++ {
				rs = append(rs, string(i+int('a')))
			}
		}
	}

	return rs
}
