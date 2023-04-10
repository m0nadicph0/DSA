package recursion

import "testing"

func Test_digitSum(t *testing.T) {

	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "single digit number",
			input:    2,
			expected: 2,
		},
		{
			name:     "two digit number",
			input:    22,
			expected: 4,
		},
		{
			name:     "three digit number",
			input:    220,
			expected: 4,
		},
		{
			name:     "four digit number",
			input:    1234,
			expected: 10,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := digitSum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected digitSum()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}

func Test_numDigits(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "single digit number",
			input:    2,
			expected: 1,
		},
		{
			name:     "single digit number zero",
			input:    0,
			expected: 1,
		},
		{
			name:     "two digit number",
			input:    22,
			expected: 2,
		},
		{
			name:     "three digit number",
			input:    220,
			expected: 3,
		},
		{
			name:     "four digit number",
			input:    1234,
			expected: 4,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := numDigits(tc.input)
			if actual != tc.expected {
				t.Errorf("expected numDigits()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}
