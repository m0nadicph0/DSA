package queue

import (
	"testing"
)

func TestLinkedListQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name         string
		queue        Queue
		input        int
		expectedSize int
	}{
		{
			name:         "enqueue to empty queue",
			queue:        initLLQueue(),
			input:        4,
			expectedSize: 1,
		},
		{
			name:         "enqueue to non-empty queue",
			queue:        initLLQueue(1, 2, 3),
			input:        4,
			expectedSize: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.queue.Enqueue(tc.input)

			actual := tc.queue.Size()
			if actual != tc.expectedSize {
				t.Errorf("expected size=%d, got=%d", tc.expectedSize, actual)
			}
		})
	}
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name         string
		queue        Queue
		expected     int
		expectedSize int
		wantError    bool
	}{
		{
			name:         "Dequeue from empty queue",
			queue:        initLLQueue(),
			expected:     -1,
			expectedSize: 0,
			wantError:    true,
		},
		{
			name:         "Dequeue from non-empty queue",
			queue:        initLLQueue(1, 2, 3),
			expected:     1,
			expectedSize: 2,
			wantError:    false,
		},
		{
			name:         "Dequeue from non-empty queue with 10 elements",
			queue:        initLLQueue(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expected:     1,
			expectedSize: 9,
			wantError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.queue.Dequeue()
			actualSize := tc.queue.Size()

			if tc.wantError && err == nil {
				t.Errorf("expected error=%v, got error=%v", tc.wantError, err)
			}
			if actual != tc.expected {
				t.Errorf("expected Dequeue to return=%d, got=%d", tc.expected, actual)
			}
			if actual != tc.expected {
				t.Errorf("expected Size to return=%d, got=%d", tc.expectedSize, actualSize)
			}
		})
	}
}

func TestLinkedListQueue_Front(t *testing.T) {
	testCases := []struct {
		name     string
		queue    Queue
		expected int
		wantErr  bool
	}{
		{
			name:     "Front on an empty queue",
			queue:    initLLQueue(),
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Front on an non-empty queue",
			queue:    initLLQueue(1, 2, 3),
			expected: 1,
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.queue.Front()

			if tc.wantErr && err == nil {
				t.Errorf("wanted error=%v, got error=%v", tc.wantErr, err)
			}

			if actual != tc.expected {
				t.Errorf("wanted Front=%d, got=%d", tc.expected, actual)
			}
		})
	}
}

func TestLinkedListQueue_IsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		queue    Queue
		expected bool
	}{
		{
			name:     "IsEmpty on an empty queue",
			queue:    initLLQueue(),
			expected: true,
		},
		{
			name:     "IsEmpty on an non-empty queue",
			queue:    initLLQueue(1, 2, 3),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.queue.IsEmpty()

			if actual != tc.expected {
				t.Errorf("wanted IsEmpty=%t, got=%t", tc.expected, actual)
			}
		})
	}
}

func TestLinkedListQueue_Rear(t *testing.T) {
	testCases := []struct {
		name     string
		queue    Queue
		expected int
		wantErr  bool
	}{
		{
			name:     "Rear on an empty queue",
			queue:    initLLQueue(),
			expected: -1,
			wantErr:  true,
		},
		{
			name:     "Rear on an non-empty queue",
			queue:    initLLQueue(4, 3, 2, 1),
			expected: 1,
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.queue.Rear()

			if tc.wantErr && err == nil {
				t.Errorf("wanted error=%v, got error=%v", tc.wantErr, err)
			}

			if actual != tc.expected {
				t.Errorf("wanted Rear=%d, got=%d", tc.expected, actual)
			}
		})
	}
}

func TestLinkedListQueue_Size(t *testing.T) {
	testCases := []struct {
		name     string
		queue    Queue
		expected int
	}{
		{
			name:     "Size on an empty queue",
			queue:    initLLQueue(),
			expected: 0,
		},
		{
			name:     "Size on an non-empty queue",
			queue:    initLLQueue(1, 2, 3),
			expected: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.queue.Size()

			if actual != tc.expected {
				t.Errorf("wanted IsEmpty=%d, got=%d", tc.expected, actual)
			}
		})
	}
}

func TestLinkedListQueue_EnqueueDequeue(t *testing.T) {
	q := NewLinkedListQueue()

	for i := 0; i < 10; i++ {
		q.Enqueue(i + 1)
	}

	front, _ := q.Front()
	rear, _ := q.Rear()
	if front != 1 && rear != 10 {
		t.Errorf("wanted front=%d and rear=%d, but got front=%d and rear =%d", 1, 10, front, rear)
	}

	if q.Size() != 10 {
		t.Errorf("wanted size=%d, but got=%d", 10, q.Size())
	}

	for !q.IsEmpty() {
		_, _ = q.Dequeue()
	}

	if q.Size() != 0 {
		t.Errorf("wanted size=%d, but got=%d", 0, q.Size())
	}

	q.Enqueue(11)
	q.Enqueue(12)

	value, _ := q.Dequeue()
	if value != 11 {
		t.Errorf("wanted dequeue=%d, but got=%d", 11, value)
	}
	value, _ = q.Dequeue()

	if value != 12 {
		t.Errorf("wanted dequeue=%d, but got=%d", 12, value)
	}
}

func initLLQueue(nums ...int) Queue {
	queue := NewLinkedListQueue()
	for _, num := range nums {
		queue.Enqueue(num)
	}
	return queue
}
