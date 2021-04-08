package issue

import (
	"strings"
	"unicode/utf8"
)

func align(words []string, length int) string {
	wordsLen := len(words)

	if wordsLen == 1 {
		wl := utf8.RuneCountInString(words[0])
		if wl >= length {
			return words[0]
		}

		return words[0] + strings.Repeat(" ", length-wl)
	}

	wordsSum := utf8.RuneCountInString(strings.Join(words, ""))
	minSpaces := wordsLen - 1
	freeSpace := length - wordsSum - minSpaces

	if freeSpace <= 0 {
		return strings.Join(words, " ")
	}

	partLen := freeSpace / wordsLen
	stringBuilder := strings.Builder{}
	tailings := freeSpace - partLen*minSpaces

	for i := range words {
		stringBuilder.WriteString(words[i])

		if i != (wordsLen - 1) {
			stringBuilder.WriteString(" ")
			s := partLen - utf8.RuneCountInString(words[i]) + 1

			if s > 0 {
				stringBuilder.WriteString(strings.Repeat(" ", s))
			}

			if i < tailings {
				stringBuilder.WriteString(" ")
			}
		}
	}

	return stringBuilder.String()
}
