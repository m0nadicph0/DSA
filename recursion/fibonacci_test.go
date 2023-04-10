package recursion

import "testing"

func Test_fibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "first fibonacci number",
			input:    0,
			expected: 0,
		},
		{
			name:     "second fibonacci number",
			input:    1,
			expected: 1,
		},
		{
			name:     "third fibonacci number",
			input:    2,
			expected: 1,
		},
		{
			name:     "fourth fibonacci number",
			input:    3,
			expected: 2,
		},
		{
			name:     "fifth fibonacci number",
			input:    4,
			expected: 3,
		},
		{
			name:     "sixth fibonacci number",
			input:    5,
			expected: 5,
		},
		{
			name:     "seventh fibonacci number",
			input:    6,
			expected: 8,
		},
		{
			name:     "eighth fibonacci number",
			input:    7,
			expected: 13,
		},
		{
			name:     "ninth fibonacci number",
			input:    8,
			expected: 21,
		},
		{
			name:     "tenth fibonacci number",
			input:    9,
			expected: 34,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := fibonacci(tc.input)
			if actual != tc.expected {
				t.Errorf("fibonacci() = %v, want %v", actual, tc.expected)
			}
		})
	}
}
