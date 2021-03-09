package issue_test

import (
	"gnotes/algorithms/issue"
	"testing"
)

func TestReverse(t *testing.T) {
	text := "some text."
	want := ".txet emos"
	reverse := issue.Reverse(text)

	if reverse != want {
		t.Fatalf("Result expected '%s', got '%s'", want, reverse)
	}
}
