package utils

import (
	"queue/queue"
	"testing"
)

func TestReverseQueue(t *testing.T) {
	q := queue.NewSliceQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	ReverseQueue(q)

	if q.IsEmpty() {
		t.Errorf("expected IsEmpty=false, but got IsEmpty=%t", q.IsEmpty())
	}

	value, _ := q.Dequeue()
	if value != 3 {
		t.Errorf("expected Dequeue=3, but got=%d", value)
	}
	value, _ = q.Dequeue()
	if value != 2 {
		t.Errorf("expected Dequeue=2, but got=%d", value)
	}

	value, _ = q.Dequeue()
	if value != 1 {
		t.Errorf("expected Dequeue=1, but got=%d", value)
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("expected error, but got=%v", err)
	}
}

func TestReverseQueueEmpty(t *testing.T) {
	q := queue.NewSliceQueue()

	ReverseQueue(q)

	if !q.IsEmpty() {
		t.Errorf("expected IsEmpty=true, but got IsEmpty=%t", q.IsEmpty())
	}
}

func TestReverseQueueRec(t *testing.T) {
	q := queue.NewSliceQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	ReverseQueueRec(q)

	if q.IsEmpty() {
		t.Errorf("expected IsEmpty=false, but got IsEmpty=%t", q.IsEmpty())
	}

	value, _ := q.Dequeue()
	if value != 3 {
		t.Errorf("expected Dequeue=3, but got=%d", value)
	}
	value, _ = q.Dequeue()
	if value != 2 {
		t.Errorf("expected Dequeue=2, but got=%d", value)
	}

	value, _ = q.Dequeue()
	if value != 1 {
		t.Errorf("expected Dequeue=1, but got=%d", value)
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("expected error, but got=%v", err)
	}
}
