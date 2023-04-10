package recursion

import "testing"

func Test_gcd(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"case 1", 1, 1, 1},
		{"case 2", 2, 3, 1},
		{"case 3", 4, 6, 2},
		{"case 4", 9, 12, 3},
		{"case 5", 17, 23, 1},
		{"case 6", 36, 48, 12},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := gcd(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("gcd(%d, %d) = %d; expected %d", tc.a, tc.b, actual, tc.expected)
			}
		})
	}
}

func Test_lcm(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"case 1", 1, 1, 1},
		{"case 2", 2, 3, 6},
		{"case 3", 4, 6, 12},
		{"case 4", 9, 12, 36},
		{"case 5", 17, 23, 391},
		{"case 6", 36, 48, 144},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lcm(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("lcm(%d, %d) = %d; expected %d", tc.a, tc.b, actual, tc.expected)
			}
		})
	}
}
