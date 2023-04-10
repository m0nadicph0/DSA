package recursion

import "testing"

func Test_factorial(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "factorial of 0 is 1",
			input:    0,
			expected: 1,
		},
		{
			name:     "factorial of 1 is 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "factorial of 2 is 2",
			input:    2,
			expected: 2,
		},
		{
			name:     "factorial of 5 is 120",
			input:    5,
			expected: 120,
		},
		{
			name:     "The factorial of 10 is 3,628,800",
			input:    10,
			expected: 3628800,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := factorial(tc.input)
			if actual != tc.expected {
				t.Errorf("factorial() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}
