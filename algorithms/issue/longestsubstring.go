package issue

import "math"

// LongestSubstring from two strings
func LongestSubstring(a, b string) int {
	table := make(map[int]map[int]int)
	longest := 0

	for i := range a {
		if _, ok := table[i]; !ok {
			table[i] = make(map[int]int)
		}

		for j := range b {
			if a[i] == b[j] {
				prev := 0

				if p, ok := table[i-1][j-1]; ok {
					prev = p
				}

				table[i][j] = prev + 1
			} else { // letters not match
				prevA := 0
				prevB := 0

				if p, ok := table[i-1][j]; ok {
					prevA = p
				}

				if p, ok := table[i][j-1]; ok {
					prevB = p
				}

				table[i][j] = int(math.Max(float64(prevA), float64(prevB)))
			}

			if table[i][j] > longest {
				longest = table[i][j]
			}
		}
	}

	return longest
}
