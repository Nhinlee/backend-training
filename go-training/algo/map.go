package algo

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if idx, found := m[diff]; found {
			return []int{m[diff], idx}
		}

		m[num] = i
	}

	return nil
}

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	// Input: strs = ["eat","tea","tan","ate","nat","bat"]
	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	m := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}

		m[count] = append(m[count], str)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// https://leetcode.com/problems/subarray-sum-equals-k/description/
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // Initialize with sum 0 having one count

	for _, num := range nums {
		sum += num
		if _, exists := prefixSum[sum-k]; exists {
			count += prefixSum[sum-k]
		}
		prefixSum[sum]++
	}

	return count
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	longest := 0
	start := 0

	for i, c := range s {
		if index, found := charIndex[c]; found && index >= start {
			start = index + 1
		}
		charIndex[c] = i
		if i-start+1 > longest {
			longest = i - start + 1
		}
	}

	return longest
}
