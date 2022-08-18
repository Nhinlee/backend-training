package solutions

import (
	"sort"
)

func minSetSize(arr []int) int {
	flags := map[int]int{}
	for _, a := range arr {
		flags[a]++
	}

	values := make([]int, len(flags))
	i := 0
	for _, e := range flags {
		values[i] = e
		i++
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	rs, count := 0, 0
	for _, v := range values {
		count += v
		rs++
		if count >= len(arr)/2 {
			break
		}
	}

	return rs
}
