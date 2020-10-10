package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items)-1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	stack := Stack{}
	fmt.Println(stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
