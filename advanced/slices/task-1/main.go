package main

import (
	"fmt"
)

// type Stack[T any] struct {
// items []T
// }

// Реалізуй методи: Push(item T), Pop() (T, bool), Peek() (T, bool), IsEmpty() bool

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	removedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return removedItem, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !stack.IsEmpty() {
		if val, ok := stack.Pop(); ok {
			fmt.Println("Popped:", val)
		}
	}

}
