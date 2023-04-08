package queue

type Queue interface {
	Enqueue(value any)
	Dequeue() (any, error)
	Front() (any, error)
	Rear() (any, error)
	Size() int
	IsEmpty() bool
}
