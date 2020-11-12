package stacks

import (
	"errors"
	"strings"

	"algos.com/main/linkedlist"
)

// StackUsingList will create a stack data structure using
// a singly linked list as the underlying data type
type StackUsingList struct {
	list linkedlist.LinkedList
	size int
}

// CreateStackUsingList will create and return a new stack using a linked list
func CreateStackUsingList(size int) StackUsingList {
	list := linkedlist.CreateLinkedList()
	stack := StackUsingList{
		list: list,
		size: size,
	}
	return stack
}

// Push will push new elements to the stack.
//
// If the stack size exceeds, it will panic with a stack overflow error
func (stack *StackUsingList) Push(data int) {
	if stack.IsFull() {
		panic(errors.New("Stack overflow error"))
	}
	stack.list.InsertFirst(data)
}

// Pop will remove elements from the stack.
//
// If the stack is empty, it will panic with a stack underflow error
func (stack *StackUsingList) Pop() {
	error := stack.list.DeleteFirst()
	if error != nil {
		panic(error)
	}
}

// IsEmpty will return true if the stack is empty and false otherwise
func (stack *StackUsingList) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// IsFull will return true if the stack is full and false otherwise
func (stack *StackUsingList) IsFull() bool {
	return stack.list.Count() == stack.size
}

// StackTop will return true the top most element, panics if the stack is empty
func (stack *StackUsingList) StackTop() int {
	return stack.list.FindFirst()
}

// Peek will return the element corresponding to the given index
//
// panics if the index is out of range
func (stack *StackUsingList) Peek(index int) int {
	if stack.IsEmpty() || index < 0 || index > stack.list.Count() {
		panic(errors.New("Index out of bound error"))
	}
	value, _ := stack.list.FindByIndex(index)
	return value
}

// String method will return a string representation of the stack
func (stack *StackUsingList) String() string {
	if stack.IsEmpty() {
		return "{}"
	}

	listRepr := stack.list.String()
	elements := strings.Split(listRepr, "-> ")
	count := len(elements)

	var sb strings.Builder

	for i := 0; i < count; i++ {
		sb.WriteString(elements[i])
		sb.WriteByte('\n')
		if i != count-1 {
			sb.WriteRune('↓')
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}
