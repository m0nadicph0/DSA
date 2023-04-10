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

func Test_power(t *testing.T) {

	tests := []struct {
		name           string
		base, exponent int
		want           int
	}{
		{"0^0 edge case", 0, 0, 1},
		{"base to 0 power", 2, 0, 1},
		{"0 to non-zero power", 0, 3, 0},
		{"any number to the power of 0 is 1", 1, 100, 1},
		{"negative base and odd exponent", -2, 3, -8},
		{"negative base and even exponent", -2, 4, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := power(tt.base, tt.exponent)
			if got != tt.want {
				t.Errorf("power(%d, %d) = %d; want %d", tt.base, tt.exponent, got, tt.want)
			}
		})
	}
}
