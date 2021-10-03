package main

type Stack []int

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Pop removes and return top element of stack.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}
