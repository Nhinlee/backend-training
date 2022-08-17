package solutions

func firstUniqChar(s string) int {
	var flags [26]int
	for _, v := range s {
		flags[v-'a']++
	}

	for i, v := range s {
		if flags[v-'a'] == 1 {
			return i
		}
	}

	return -1
}
