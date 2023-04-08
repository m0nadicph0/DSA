package queue

import "errors"

type SliceQueue struct {
	buffer []any
}

func NewSliceQueue() Queue {
	return &SliceQueue{
		buffer: make([]any, 0),
	}
}

func (q *SliceQueue) Enqueue(value any) {
	q.buffer = append(q.buffer, value)
}

func (q *SliceQueue) Dequeue() (any, error) {
	if len(q.buffer) == 0 {
		return -1, errors.New("queue is empty")
	}
	item := q.buffer[0]
	q.buffer = q.buffer[1:]
	return item, nil
}

func (q *SliceQueue) Front() (any, error) {
	if len(q.buffer) == 0 {
		return -1, errors.New("queue is empty")
	}
	return q.buffer[0], nil
}

func (q *SliceQueue) Rear() (any, error) {
	if len(q.buffer) == 0 {
		return -1, errors.New("queue is empty")
	}
	return q.buffer[len(q.buffer)-1], nil
}

func (q *SliceQueue) Size() int {
	return len(q.buffer)
}

func (q *SliceQueue) IsEmpty() bool {
	return len(q.buffer) == 0
}
