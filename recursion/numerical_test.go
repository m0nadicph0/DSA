package recursion

import (
	"reflect"
	"testing"
)

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

/*
Test the first triangular number: assert triangular_number(1) == 1
Test the fifth triangular number: assert triangular_number(5) == 15
Test the tenth triangular number: assert triangular_number(10) == 55
Test the 100th triangular number: assert triangular_number(100) == 5050
Test the function with a negative input: assert triangular_number(-1) == None

*/
func Test_triangularNth(t *testing.T) {

	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "first triangular number",
			n:    1,
			want: 1,
		},
		{
			name: "fifth triangular number",
			n:    5,
			want: 15,
		},
		{
			name: "tenth triangular number",
			n:    10,
			want: 55,
		},
		{
			name: "100th triangular number",
			n:    100,
			want: 5050,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := triangularNth(tc.n)
			if got != tc.want {
				t.Errorf("triangularNth(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_sumNatural(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 0,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 1,
		},
		{
			name: "sum of first 5 natural numbers",
			n:    5,
			want: 15,
		},
		{
			name: "sum of first 10 natural numbers",
			n:    10,
			want: 55,
		},
		{
			name: "sum of first 100 natural numbers",
			n:    100,
			want: 5050,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumNatural(tc.n)
			if got != tc.want {
				t.Errorf("sumNatural(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_sumOfFirstNEven(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 0,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 2,
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: 42,
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: 30,
		},
		{
			name: "When n is 10",
			n:    10,
			want: 110,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumOfFirstNEven(tc.n)
			if got != tc.want {
				t.Errorf("sumOfFirstNEven(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_sumOfFirstNOdd(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 0,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 1,
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: 36,
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: 25,
		},
		{
			name: "When n is 10",
			n:    10,
			want: 100,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := sumOfFirstNOdd(tc.n)
			if got != tc.want {
				t.Errorf("sumOfFirstNOdd(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_forEachUptoN(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: []int{},
		},
		{
			name: "for n equals 1",
			n:    1,
			want: []int{1},
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "When n is 10",
			n:    10,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := make([]int, 0)
			forEachUptoN(tc.n, func(i int) {
				got = append(got, i)
			})
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("forEachUptoN(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_forEachFromNDownTo1(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: []int{},
		},
		{
			name: "for n equals 1",
			n:    1,
			want: []int{1},
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "When n is 10",
			n:    10,
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := make([]int, 0)
			forEachFromNDownTo1(tc.n, func(i int) {
				got = append(got, i)
			})
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("forEachFromNDownTo1(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_productNatural(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 1,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 1,
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: 720,
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: 120,
		},
		{
			name: "When n is 10",
			n:    10,
			want: 3628800,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := productNatural(tc.n)

			if got != tc.want {
				t.Errorf("productNatural(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_productEvenNatural(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 1,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 2,
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: 46080,
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: 3840,
		},
		{
			name: "When n is 10",
			n:    10,
			want: 3715891200,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := productEvenNatural(tc.n)

			if got != tc.want {
				t.Errorf("productEvenNatural(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_productOddNatural(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 1,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 1,
		},
		{
			name: "When n is a positive even integer",
			n:    6,
			want: 10395,
		},
		{
			name: "When n is a positive odd integer",
			n:    5,
			want: 945,
		},
		{
			name: "When n is 10",
			n:    10,
			want: 654729075,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := productOddNatural(tc.n)

			if got != tc.want {
				t.Errorf("productOddNatural(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func Test_sumDigitsBinary(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "for n equals 0",
			n:    0,
			want: 0,
		},
		{
			name: "for n equals 1",
			n:    1,
			want: 1,
		},
		{
			name: "When n 10",
			n:    10,
			want: 1,
		},
		{
			name: "When n 11",
			n:    11,
			want: 2,
		},
		{
			name: "When n is 1010",
			n:    1010,
			want: 2,
		},
		{
			name: "When n is 1101",
			n:    1101,
			want: 3,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			got := sumDigitsBinary(tc.n)

			if got != tc.want {
				t.Errorf("sumDigitsBinary(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
