package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{3, 5, 4, 1, 2}
	want := []int{1, 2, 3, 4, 5}

	result := SelectionSort(list)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Wrong result %v, result must be %v", result, want)
	}
}
