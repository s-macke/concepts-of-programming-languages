package stackgeneric

import "fmt"

package main

import (
    "fmt"
)

// Stack is a generic LIFO container for untyped object.
type Stack[T any] struct {
	data []T
}

// NewStack constructs an empty stack.
func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

// Push pushes a value on the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop pops a value from the stack. It returns an error if the stack is empty.
func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		panic("can not pop: empty stack")
	}
	var result = s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return result
}

// Size returns the count of elements in the Stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// Get returns the n-th element in the Stack
func (s *Stack[T]) Get(idx int) T {
	return s.data[idx]
}

func main() {
	s := NewStack[int]()
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
