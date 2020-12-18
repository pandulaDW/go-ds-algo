package trees

import "algos.com/main/queues"

// Stack represents a simple stack implementation
type Stack struct {
	data queues.LinkedList
}

// CreateSimpleStack creates a new stack
func CreateSimpleStack() *Stack {
	list := queues.CreateLinkedList()
	return &Stack{*list}
}

// Push will push elements to stack
func (s *Stack) Push(item interface{}) {
	s.data.InsertFirst(item)
}

// Pop will pop the elements from the stack
func (s *Stack) Pop() interface{} {
	poppedItem := s.data.First()
	s.data.Pull()
	return poppedItem
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.data.Size() == 0
}
