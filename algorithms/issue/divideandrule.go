package issue

// just issues from book, chapter "divide and rule"

// Sum recursively calculates the sum of list
func Sum(list []int) int {
	switch len(list) {
	case 0:
		return 0
	case 1:
		return list[0]
	}

	return list[0] + Sum(list[1:])
}

// Count recursively calculates the count of list
func Count(list []int) int {
	// its stupid, but how else...
	if len(list) == 0 {
		return 0
	}

	return Count(list[1:]) + 1
}

// Max recursively calculates the max of list
func Max(list []int) int {
	switch len(list) {
	case 0:
		return 0
	case 1:
		return list[0]
	}

	max := Max(list[1:])

	if list[0] > max {
		max = list[0]
	}

	return max
}

// BinarySearch recursively
func BinarySearch(list []int, item int) (mid int, ok bool) {
	switch len(list) {
	case 0:
		return 0, false
	case 1:
		return 0, list[0] == item
	}

	high := len(list) - 1
	mid = high / 2
	guess := list[mid]

	switch {
	case guess == item:
		return mid, true
	case mid == 1:
		return 0, false
	case guess < item:
		r, ok := BinarySearch(list[mid:], item)
		if !ok {
			mid = 0
		}

		return mid + r, ok
	}

	return BinarySearch(list[:mid+1], item)
}
