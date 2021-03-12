package issue_test

import (
	"gnotes/algorithms/issue"
	"reflect"
	"testing"
)

func TestNumberPipeline(t *testing.T) {
	result := issue.NumberPipeline(10)
	want := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	if !reflect.DeepEqual(result, want) {
		t.Fatalf("Result must be: %v, got: %v", want, result)
	}
}
