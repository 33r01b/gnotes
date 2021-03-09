package sort_test

import (
	"gnotes/algorithms/sort"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{5, 3, 2, 4, 1}
	want := []int{1, 2, 3, 4, 5}

	result := sort.QuickSort(list)

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Wrong result %v, result must be %v", result, want)
	}
}
