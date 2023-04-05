package stack

type Stack interface {
	Push(value int) error
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	IsFull() bool
	Size() int
}
