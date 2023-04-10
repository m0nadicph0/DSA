package recursion

import "testing"

func Test_reverse(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "non-empty string",
			input:    "abc",
			expected: "cba",
		},
		{
			name:     "palindrome",
			input:    "madam",
			expected: "madam",
		},
		{
			name:     "multi-word palindrome",
			input:    "able was i ere i saw elba",
			expected: "able was i ere i saw elba",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := reverse(tc.input)
			if actual != tc.expected {
				t.Errorf("expected reverse()=%v, got=%v", tc.expected, actual)
			}
		})
	}
}
