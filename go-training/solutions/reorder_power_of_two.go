package solutions

func reorderedPowerOf2(n int) bool {
	count := countDigit(n)
	p := 1

	for i := 0; i <= 31; i++ {
		if count == countDigit(p) {
			return true
		}
		p <<= 1
	}
	return false
}

func countDigit(num int) [10]int {
	count := [10]int{}
	for num > 0 {
		count[num%10]++
		num /= 10
	}

	return count
}
