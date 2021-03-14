package issue_test

import (
	"gnotes/algorithms/issue"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	a := "blue"
	b := "clues"

	result := issue.LongestSubstring(a, b)
	want := len("lue")

	if result != want {
		t.Fatalf("Result must be: %d, got %d", want, result)
	}
}
