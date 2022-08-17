package solutions

// https://leetcode.com/problems/find-and-replace-pattern/

func FindAndReplacePattern(words []string, pattern string) []string {
	rs := make([]string, 0)
	for _, word := range words {
		if match(word, pattern) {
			rs = append(rs, word)
		}
	}

	return rs
}

func match(word, pattern string) bool {
	wordMap := make(map[byte]byte)
	patternMap := make(map[byte]byte)

	for i := 0; i < len(word); i++ {
		if wordMap[word[i]] != 0 && wordMap[word[i]] != pattern[i] {
			return false
		}
		if patternMap[pattern[i]] != 0 && patternMap[pattern[i]] != word[i] {
			return false
		}

		wordMap[word[i]] = pattern[i]
		patternMap[pattern[i]] = word[i]
	}
	return true
}
