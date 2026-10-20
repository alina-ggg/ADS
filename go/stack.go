package main

import (
	"fmt"
)

type StackNode struct {
	data string
	next *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Push(value string) {
	s.top = &StackNode{data: value, next: s.top}
}

func (s *Stack) Pop() string {
	if s.top == nil {
		return ""
	}
	val := s.top.data
	s.top = s.top.next
	return val
}

func (s *Stack) Print() {
	curr := s.top
	for curr != nil {
		fmt.Printf("%s ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
