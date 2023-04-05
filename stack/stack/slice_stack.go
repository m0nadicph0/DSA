package stack

import (
	"errors"
)

type Any interface{}

type SliceStack struct {
	buffer  []Any
	maxSize int
}

func NewSliceStack(size int) Stack {
	return &SliceStack{
		buffer:  make([]Any, 0),
		maxSize: size,
	}
}

func (s *SliceStack) Size() int {
	return len(s.buffer)
}

func (s *SliceStack) Push(value Any) error {
	if len(s.buffer) >= s.maxSize {
		return errors.New("stack is full")
	}
	s.buffer = append(s.buffer, value)
	return nil
}

func (s *SliceStack) Pop() (Any, error) {
	if len(s.buffer) == 0 {
		return -1, errors.New("stack is empty")
	}
	topItem := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]
	return topItem, nil
}

func (s *SliceStack) Peek() (Any, error) {
	if len(s.buffer) == 0 {
		return -1, errors.New("stack is empty")
	}
	return s.buffer[len(s.buffer)-1], nil
}

func (s *SliceStack) IsEmpty() bool {
	return len(s.buffer) == 0
}

func (s *SliceStack) IsFull() bool {
	return len(s.buffer) == s.maxSize
}
