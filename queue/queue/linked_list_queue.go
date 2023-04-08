package queue

import (
	"errors"
)

type node struct {
	value any
	next  *node
}

type LinkedListQueue struct {
	front *node
	rear  *node
	size  int
}

func NewLinkedListQueue() Queue {
	return &LinkedListQueue{
		front: nil,
		rear:  nil,
	}
}

func (q *LinkedListQueue) Enqueue(value any) {
	newNode := &node{
		value: value,
		next:  nil,
	}

	if q.rear == nil {
		q.rear = newNode
		q.front = newNode
	} else {
		q.rear.next = newNode
		q.rear = q.rear.next
	}
	q.size++
}

func (q *LinkedListQueue) Dequeue() (any, error) {
	if q.front == nil {
		return -1, errors.New("queue is empty")
	} else {
		value := q.front.value
		q.front = q.front.next
		if q.front == nil {
			q.rear = nil
		}
		q.size--
		return value, nil
	}
}

func (q *LinkedListQueue) Front() (any, error) {
	if q.front == nil {
		return -1, errors.New("queue is empty")
	}
	return q.front.value, nil
}

func (q *LinkedListQueue) Rear() (any, error) {
	if q.rear == nil {
		return -1, errors.New("queue is empty")
	}
	return q.rear.value, nil
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}
