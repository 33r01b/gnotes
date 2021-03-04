package sort

// SelectionSort realization
func SelectionSort(list []int) []int {
	newList := make([]int, len(list))

	for i := range list {
		smallest := findSmallest(list)
		newList[i] = list[smallest]
		list = append(list[:smallest], list[smallest+1:]...)
	}

	return newList
}

func findSmallest(list []int) (smallestIndex int) {
	smallest := list[smallestIndex]

	// start from 1, to drop first element
	// because smallest == list[0]
	for i := 1; i < len(list); i++ {
		if list[i] < smallest {
			smallest = list[i]
			smallestIndex = i
		}
	}

	return
}
