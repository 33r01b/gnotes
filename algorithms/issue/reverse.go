package issue

// Reverse text
// O(n)
func Reverse(text string) string {
	reversed := make([]byte, len(text))

	j := 0
	for i := len(text) - 1; i >= 0; i-- {
		reversed[j] = text[i]
		j++
	}

	return string(reversed)
}
