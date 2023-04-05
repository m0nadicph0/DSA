package main

import (
	"fmt"
	"stack/stack"
)

func pushItems(s stack.Stack, items []int) {
	for _, item := range items {
		err := s.Push(item)
		if err != nil {
			continue
		}
	}
}

func popItems(s stack.Stack) []int {
	poppedItems := make([]int, 0)
	for !s.IsEmpty() {
		item, _ := s.Pop()
		poppedItems = append(poppedItems, item)
	}
	return poppedItems
}

func main() {
	s := stack.NewSliceStack(5)
	pushItems(s, []int{1, 2, 3, 4, 5, 6})
	fmt.Println("Stack full:", s.IsFull(), "Stack empty", s.IsEmpty())
	fmt.Println(popItems(s))
	fmt.Println("Stack full:", s.IsFull(), "Stack empty", s.IsEmpty())
}
