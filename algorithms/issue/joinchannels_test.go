package issue_test

import (
	"gnotes/algorithms/issue"
	"reflect"
	"sort"
	"testing"
)

func TestJoinChannels(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{300, 200, 100} {
			c <- num
		}
		close(c)
	}()

	result := make([]int, 0)

	for num := range issue.JoinChannels(a, b, c) {
		result = append(result, num)
	}

	sort.Ints(result)

	want := []int{1, 2, 3, 10, 20, 30, 100, 200, 300}

	if !reflect.DeepEqual(want, result) {
		t.Fatalf("Result must be %v, got %v", want, result)
	}
}
