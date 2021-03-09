package issue

import "testing"

func TestReverse(t *testing.T) {
	text := "some text."
	want := ".txet emos"
	reverse := Reverse(text)

	if reverse != want {
		t.Fatalf("Result expected '%s', got '%s'", want, reverse)
	}
}
