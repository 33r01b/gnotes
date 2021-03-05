package issue

// Sum recursively calculates the sum of list
// just issue from book, chapter "divide and rule"
func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + Sum(list[1:])
}
