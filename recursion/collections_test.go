package recursion

import "testing"

func Test_sum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "non-empty array with single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "non-empty array with multiple element",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "non-empty array with multiple same element",
			input:    []int{2, 2, 2, 2, 2},
			expected: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected sum=%d, but got=%d", tc.expected, actual)
			}
		})
	}
}

func Test_product(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "non-empty array with single element",
			input:    []int{1},
			expected: 1,
		},
		{
			name:     "non-empty array with multiple element",
			input:    []int{1, 2, 3, 4, 5},
			expected: 120,
		},
		{
			name:     "non-empty array with multiple same element",
			input:    []int{2, 2, 2, 2, 2},
			expected: 32,
		},
		{
			name:     "non-empty array with multiple different element",
			input:    []int{2, 8, 2, 1, 1},
			expected: 32,
		},
		{
			name:     "non-empty array with multiple element one of them being zero",
			input:    []int{0, 1, 2, 3, 4},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := product(tc.input)
			if actual != tc.expected {
				t.Errorf("expected product=%d, but got=%d", tc.expected, actual)
			}
		})
	}
}

func Test_minimum(t *testing.T) {

	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "non-empty array of size 1",
			input:    []int{2},
			expected: 2,
		},
		{
			name:     "non-empty array",
			input:    []int{2, 4, 6, 8},
			expected: 2,
		},
		{
			name:     "non-empty array variant 1",
			input:    []int{10, 2, 8, 4, 1, 1, 22, 3004},
			expected: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := minimum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected minimum()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}

func Test_maximum(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "non-empty array of size 1",
			input:    []int{2},
			expected: 2,
		},
		{
			name:     "non-empty array",
			input:    []int{2, 4, 6, 8},
			expected: 8,
		},
		{
			name:     "non-empty array variant 1",
			input:    []int{10, 2, 8, 4, 1, 1, 22, 3004},
			expected: 3004,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := maximum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected maximum()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}

func Test_isSortedAscending(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "non-empty array of size 1",
			input:    []int{2},
			expected: true,
		},
		{
			name:     "non-empty array",
			input:    []int{2, 4, 6, 8},
			expected: true,
		},
		{
			name:     "non-empty array variant 1",
			input:    []int{10, 2, 8, 4, 1, 1, 22, 3004},
			expected: false,
		},
		{
			name:     "non-empty array variant 2",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "non-empty array variant 3",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: true,
		},
		{
			name:     "non-empty array variant 4",
			input:    []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "non-empty array variant 5",
			input:    []int{1, 1, 2, 2, 2},
			expected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSortedAscending(tc.input)
			if actual != tc.expected {
				t.Errorf("expected isSortedAscending()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}

func Test_isSortedDescending(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "non-empty array of size 1",
			input:    []int{2},
			expected: true,
		},
		{
			name:     "non-empty array",
			input:    []int{2, 4, 6, 8},
			expected: false,
		},
		{
			name:     "non-empty array variant 1",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: true,
		},
		{
			name:     "non-empty array variant 2",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: false,
		},
		{
			name:     "non-empty array variant 3",
			input:    []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "non-empty array variant 4",
			input:    []int{2, 2, 2, 1, 1},
			expected: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isSortedDescending(tc.input)
			if actual != tc.expected {
				t.Errorf("expected isSortedDescending()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}

func Test_search(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		target   int
		expected bool
	}{
		{
			name:     "empty array",
			input:    []int{},
			target:   2,
			expected: false,
		},
		{
			name:     "non-empty array of size 1",
			input:    []int{2},
			target:   2,
			expected: true,
		},
		{
			name:     "non-empty array",
			input:    []int{2, 4, 6, 8},
			target:   8,
			expected: true,
		},
		{
			name:     "non-empty array variant 1",
			input:    []int{6, 5, 4, 3, 2, 1},
			target:   0,
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := search(tc.input, tc.target)
			if actual != tc.expected {
				t.Errorf("expected search()=%v, but got=%v", tc.expected, actual)
			}
		})
	}
}
