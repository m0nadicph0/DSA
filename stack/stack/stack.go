package stack

type Stack interface {
	Push(value Any) error
	Pop() (Any, error)
	Peek() (Any, error)
	IsEmpty() bool
	IsFull() bool
	Size() int
}
