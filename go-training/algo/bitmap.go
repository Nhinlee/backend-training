package algo

func numSteps(s string) int {
	N := len(s)

	operations := 0
	carry := 0
	for i := N - 1; i > 0; i-- {
		if (int(s[i]-'0')+carry)%2 == 1 {
			operations += 2
			carry = 1
		} else {
			operations++
		}
	}

	return operations + carry
}

// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/?envType=daily-question&envId=2024-05-30
func countTriplets(arr []int) int {
	// Insert 0 at the beginning of the array
	arr = append([]int{0}, arr...)
	n := len(arr)
	rs := 0

	// Compute prefix XOR
	for i := 1; i < n; i++ {
		arr[i] ^= arr[i-1]
	}

	// Find triplets
	for i := 0; i < n; i++ {
		for k := i + 1; k < n; k++ {
			if arr[i] == arr[k] {
				rs += k - i - 1
			}
		}
	}

	return rs
}
