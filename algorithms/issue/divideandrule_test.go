package issue

import "testing"

func TestSum(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	want := 15

	result := Sum(list)

	if result != want {
		t.Fatalf("Wrong result %d, want %d", result, want)
	}
}
