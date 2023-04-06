package stack

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "reverse empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "reverse non-empty string",
			input:    "abc",
			expected: "cba",
		},
		{
			name:     "reverse longer non-empty string",
			input:    "able was i ere i saw elba",
			expected: "able was i ere i saw elba",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReverseString(tc.input)

			if actual != tc.expected {
				t.Errorf("expected top element=%s, got=%s", tc.expected, actual)
			}
		})
	}
}

func TestIsParenthesesBalanced(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty expression",
			input:    "",
			expected: false,
		},
		{
			name:     "non-empty expression with single type of parenthesis",
			input:    "((()()()(())))",
			expected: true,
		},
		{
			name:     "non-empty expression with single type of parenthesis unbalanced",
			input:    "(((())))))))))))",
			expected: false,
		},
		{
			name:     "non-empty expression with single type of parenthesis unbalanced variation",
			input:    "(((((((((",
			expected: false,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis variation 1",
			input:    "(){}[]",
			expected: true,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis variation with spaces",
			input:    "( ) [ ] { }",
			expected: true,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis variation nested",
			input:    "( {[{[()]}]} ) [({[[[]]]})] {([{()[]{}}])}",
			expected: true,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis unbalanced",
			input:    "{[(}])",
			expected: false,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis unbalanced variation",
			input:    "{[({[(",
			expected: false,
		},
		{
			name:     "non-empty expression with multiple type of parenthesis unbalanced variation 2",
			input:    "}{][)(",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsParenthesesBalanced(tc.input)

			if actual != tc.expected {
				t.Errorf("expected top element=%t, got=%t", tc.expected, actual)
			}
		})
	}
}
