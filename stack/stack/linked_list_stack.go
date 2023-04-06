package stack

import "errors"

type Node struct {
	value Any
	next  *Node
}

type LinkedListStack struct {
	head    *Node
	maxSize int
}

func NewLinkedListStack(size int) Stack {
	return &LinkedListStack{head: nil, maxSize: size}
}

func (s *LinkedListStack) Push(value Any) error {
	if s.Size() >= s.maxSize {
		return errors.New("stack is full")
	}

	if s.head == nil {
		s.head = &Node{
			value: value,
			next:  nil,
		}
	} else {
		current := s.head
		s.head = &Node{
			value: value,
			next:  current,
		}
	}
	return nil
}

func (s *LinkedListStack) Pop() (Any, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	value := s.head.value
	s.head = s.head.next
	return value, nil
}

func (s *LinkedListStack) Peek() (Any, error) {
	if s.head == nil {
		return -1, errors.New("stack is empty")
	}

	return s.head.value, nil
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.head == nil
}

func (s *LinkedListStack) IsFull() bool {
	return s.Size() == s.maxSize
}

func (s *LinkedListStack) Size() int {
	current := s.head
	count := 0
	for current != nil {
		count++
		current = current.next
	}
	return count
}
