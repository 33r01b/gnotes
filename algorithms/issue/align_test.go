package issue

import "testing"

func TestAlign(t *testing.T) {
	testCases := []struct {
		name         string
		words        []string
		stringLength int
		expectedStr  string
	}{
		{
			name:         "case 0",
			words:        []string{"aaa"},
			stringLength: 3,
			expectedStr:  "aaa",
		},
		{
			name:         "case 1",
			words:        []string{"aaa"},
			stringLength: 4,
			expectedStr:  "aaa ",
		},
		{
			name:         "case 2",
			words:        []string{"a", "b", "c"},
			stringLength: 7,
			expectedStr:  "a  b  c",
		},
		{
			name:         "case 3",
			words:        []string{"a", "b", "c"},
			stringLength: 8,
			expectedStr:  "a   b  c",
		},
		{
			name:         "case 4",
			words:        []string{"rdst", "dds", "jkdss", "cs"},
			stringLength: 14 + 3,
			expectedStr:  "rdst dds jkdss cs",
		},
		{
			name:         "case 5",
			words:        []string{"rdst", "dds", "jkdss", "cs"},
			stringLength: 14 + 4,
			expectedStr:  "rdst  dds jkdss cs",
		},
		{
			name:         "case 6",
			words:        []string{"rdst", "dds", "jkdss", "cs"},
			stringLength: 14 + 5,
			expectedStr:  "rdst  dds  jkdss cs",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := align(tc.words, tc.stringLength)

			if result != tc.expectedStr {
				t.Fatalf("Wrong result '%s', expected '%s'", result, tc.expectedStr)
			}
		})
	}
}
