package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int

// IsEmpty checks if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop removes and return top element of stack.
func (s *stack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

// Get returns the n-th element in the Stack
func (s *stack) Get(idx int) int {
	return (*s)[len(*s)-idx-1]
}

func main() {
	var s stack
	fields := strings.Fields("2 3 5 * +")
	for _, expr := range fields {
		fmt.Println("Parse ", expr)
		switch expr {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		default:
			v, err := strconv.Atoi(expr)
			if err != nil {
				fmt.Println("Unknown expr", expr)
				return
			}
			s.Push(v)
		}
	}
	println("result:", s.Pop())
}
