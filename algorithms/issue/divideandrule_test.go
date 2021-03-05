package issue

import (
	"testing"
)

func TestSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	want := 15

	result := Sum(list)

	if result != want {
		t.Fatalf("Wrong result %d, want %d", result, want)
	}
}

func TestCount(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	want := 5

	result := Count(list)

	if result != want {
		t.Fatalf("Wrong result %d, want %d", result, want)
	}
}

func TestMax(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	want := 5

	result := Max(list)

	if result != want {
		t.Fatalf("Wrong result %d, want %d", result, want)
	}
}

func TestBinarySearch(t *testing.T) {
	list := []int{1, 12, 13, 41, 51, 63, 117, 800}

	testCases := []struct {
		name    string
		value   int
		wantKey int
	}{
		{
			name:    "case 0",
			value:   1,
			wantKey: 0,
		},
		{
			name:    "case 1",
			value:   117,
			wantKey: 6,
		},
		{
			name:    "case 2",
			value:   12,
			wantKey: 1,
		},
		{
			name:    "case 3",
			value:   13,
			wantKey: 2,
		},
		{
			name:    "case 4",
			value:   41,
			wantKey: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, ok := BinarySearch(list, tc.value)

			if result != tc.wantKey && ok {
				t.Fatalf("Wrong result %d, want %d", result, tc.wantKey)
			}
		})
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result, ok := BinarySearch(list, 10)

	if ok || result != 0 {
		t.Fatalf("Wrong result, result must be not found, got %d", result)
	}
}
