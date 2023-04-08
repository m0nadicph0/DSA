package stack

import "errors"

type Stack interface {
	Push(value any)
	Pop() (any, error)
	Peek() (any, error)
	IsEmpty() bool
	Size() int
}

func NewDynamicSliceStack() Stack {
	return &DynamicSliceStack{buffer: make([]any, 0)}
}

type DynamicSliceStack struct {
	buffer []any
}

func (s *DynamicSliceStack) Size() int {
	return len(s.buffer)
}

func (s *DynamicSliceStack) Push(value any) {
	s.buffer = append(s.buffer, value)
}

func (s *DynamicSliceStack) Pop() (any, error) {
	if len(s.buffer) == 0 {
		return -1, errors.New("stack is empty")
	}

	value := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]
	return value, nil
}

func (s *DynamicSliceStack) Peek() (any, error) {
	if len(s.buffer) == 0 {
		return -1, errors.New("stack is empty")
	}
	return s.buffer[len(s.buffer)-1], nil
}

func (s *DynamicSliceStack) IsEmpty() bool {
	return len(s.buffer) == 0
}
