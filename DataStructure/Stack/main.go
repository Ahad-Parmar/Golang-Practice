package main

import (
	"fmt"
)

type item struct {
	value interface{} 
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &item{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.Len() > 0 {
		value = s.top.value
		s.top = s.top.next
		s.size--
		return
	}

	return nil
}

func main() {
	s := new(Stack)
	
	s.Push(121)
	s.Push("one twenty one")
	s.Push(121.0)

	
	for s.Len() > 0 {
		fmt.Println(s.Pop())
	}
}
