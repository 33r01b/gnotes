package search_test

import (
	"gnotes/algorithms/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 12, 13, 41, 51, 63, 117, 800}
	wantKey := 6

	result, _ := search.BinarySearch(list, 117)

	if wantKey != result {
		t.Fatalf("Wrong result, result must be %d, got %d", wantKey, result)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result, ok := search.BinarySearch(list, 10)

	if ok {
		t.Fatalf("Wrong result, result must be not found, got %d", result)
	}
}
