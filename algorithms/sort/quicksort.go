package sort

// QuickSort realization
func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	// select middle element of list, if list is sorted
	// for other cases it does not matter
	pivot := list[(len(list)-1)/2]

	less := make([]int, 0)
	greater := make([]int, 0)

	for i := 0; i < len(list); i++ {
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
