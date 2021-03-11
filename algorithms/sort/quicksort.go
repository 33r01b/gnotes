package sort

// QuickSort realization
// O(n log n)
func QuickSort(list []int) []int {
	listLength := len(list)
	if listLength < 2 {
		return list
	}

	// select middle element of list, in case of list is already sorted
	// for other cases it does not matter
	pivot := list[(listLength-1)/2]

	less := make([]int, 0, listLength/2)
	greater := make([]int, 0, listLength/2)

	for i := 0; i < listLength; i++ {
		switch {
		case list[i] < pivot:
			less = append(less, list[i])
			break
		case list[i] > pivot:
			greater = append(greater, list[i])
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
