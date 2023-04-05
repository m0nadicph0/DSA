package stack

import "testing"

func TestSliceStack_Push(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		value    int
		wantErr  bool
	}{
		{
			name:     "Push onto non-full stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			value:    4,
			wantErr:  false,
		},
		{
			name:     "Push onto full stack",
			maxSize:  3,
			elements: []int{1, 2, 3},
			value:    4,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			err := s.Push(tc.value)
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got error=%v", tc.wantErr, err)
			}
			actual, _ := s.Peek()
			if !tc.wantErr && actual != tc.value {
				t.Errorf("expected top element=%d, got=%d", tc.value, actual)
			}
		})
	}

}

func TestSliceStack_Pop(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		expected int
		wantErr  bool
	}{
		{
			name:     "Pop from non-empty stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "Pop from empty stack",
			maxSize:  5,
			elements: []int{},
			expected: -1,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			got, err := s.Pop()
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got error=%v", tc.wantErr, err)
			}

			if got != tc.expected {
				t.Errorf("expected=%d, got=%d", tc.expected, got)
			}
		})
	}
}

func TestSliceStack_IsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		expected bool
	}{
		{
			name:     "IsEmpty of non-empty stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "IsEmpty of empty stack",
			maxSize:  5,
			elements: []int{},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			got := s.IsEmpty()
			if got != tc.expected {
				t.Errorf("expected=%t, got=%t", tc.expected, got)
			}
		})
	}
}

func TestSliceStack_IsFull(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		expected bool
	}{
		{
			name:     "IsFull of non-empty stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "IsFull of empty stack",
			maxSize:  5,
			elements: []int{},
			expected: false,
		},
		{
			name:     "IsFull of full stack",
			maxSize:  5,
			elements: []int{1, 2, 3, 4, 5},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			got := s.IsFull()
			if got != tc.expected {
				t.Errorf("expected=%t, got=%t", tc.expected, got)
			}
		})
	}
}

func TestSliceStack_Size(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		expected int
	}{
		{
			name:     "Size of non-empty stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "Size of empty stack",
			maxSize:  5,
			elements: []int{},
			expected: 0,
		},
		{
			name:     "Size of full stack",
			maxSize:  5,
			elements: []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "Size of stack after overflow",
			maxSize:  5,
			elements: []int{1, 2, 3, 4, 5, 6, 7},
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			got := s.Size()
			if got != tc.expected {
				t.Errorf("expected=%d, got=%d", tc.expected, got)
			}
		})
	}
}

func TestSliceStack_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		maxSize  int
		elements []int
		expected int
		wantErr  bool
	}{
		{
			name:     "Peek from non-empty stack",
			maxSize:  5,
			elements: []int{1, 2, 3},
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "Peek from empty stack",
			maxSize:  5,
			elements: []int{},
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Peek from full stack",
			maxSize:  5,
			elements: []int{1, 2, 3, 4, 5, 6},
			expected: 5,
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSliceStack(tc.maxSize)
			for _, elem := range tc.elements {
				_ = s.Push(elem)
			}

			got, err := s.Peek()
			if (err != nil) != tc.wantErr {
				t.Errorf("expected error=%v, got error=%v", tc.wantErr, err)
			}

			if got != tc.expected {
				t.Errorf("expected=%d, got=%d", tc.expected, got)
			}
		})
	}
}
