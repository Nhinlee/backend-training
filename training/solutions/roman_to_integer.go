package solutions

func romanToInt(s string) int {
	rs, n := 0, len(s)
	mapper := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < n-1; i++ {
		num := mapper[s[i]]
		if num < mapper[s[i+1]] {
			rs -= num
		} else {
			rs += num
		}
	}
	rs += mapper[s[n-1]]

	return rs
}
