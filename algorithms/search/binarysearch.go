package search

// BinarySearch return key of list by value
func BinarySearch(list []int, item int) (mid int, ok bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid = (low + high) / 2
		guess := list[mid]

		switch {
		case guess == item:
			return mid, true
		case guess > item:
			high = mid - 1
			break
		default:
			low = mid + 1
		}
	}

	return 0, false
}
