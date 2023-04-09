package utils

import (
	"queue/queue"
	"queue/stack"
)

func ReverseQueue(queue queue.Queue) {
	s := stack.NewDynamicSliceStack()

	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		s.Push(item)
	}

	for !s.IsEmpty() {
		item, _ := s.Pop()
		queue.Enqueue(item)
	}
}

func ReverseQueueRec(queue queue.Queue) {
	if queue.IsEmpty() {
		return
	}
	item, _ := queue.Dequeue()
	ReverseQueue(queue)
	queue.Enqueue(item)
}
