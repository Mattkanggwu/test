package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	data []interface{}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(element interface{}) {
	s.data = append(s.data, element)
}

// Pop removes and returns the element from the top of the stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(s.data) - 1
	element := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return element
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack size:", stack.Size())

	for !stack.IsEmpty() {
		element := stack.Pop()
		fmt.Println("Popped element:", element)
	}
}
