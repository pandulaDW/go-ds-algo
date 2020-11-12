package stacks

import (
	"errors"
	"strconv"
	"strings"
)

// StackUsingArray will create a stack data structure using
// arrays as the underlying data type
type StackUsingArray struct {
	data []int
	size int
	top  int
}

// CreateStackUsingArray will create and return a new stack using arrays
func CreateStackUsingArray(size int) StackUsingArray {
	stack := StackUsingArray{
		data: make([]int, 0, size),
		size: size,
		top:  -1,
	}

	return stack
}

// Push will push new elements to the stack.
//
// If the stack size exceeds, it will panic with a stack overflow error
func (stack *StackUsingArray) Push(data int) {
	if stack.IsFull() {
		panic(errors.New("Stack overflow error"))
	}
	stack.data = append(stack.data, data)
	stack.top++
}

// Pop will remove elements from the stack.
//
// If the stack is empty, it will panic with a stack underflow error
func (stack *StackUsingArray) Pop() {
	if stack.IsEmpty() {
		panic(errors.New("Stack underflow error"))
	}
	stack.data = stack.data[:stack.top]
	stack.top--
}

// Peek will return the element corresponding to the given index
//
// panics if the index is out of range
func (stack *StackUsingArray) Peek(index int) int {
	if stack.IsEmpty() || index < 0 || index > stack.top {
		panic(errors.New("Index out of bound error"))
	}
	return stack.data[stack.top-index]
}

// IsEmpty will return true if the stack is empty and false otherwise
func (stack *StackUsingArray) IsEmpty() bool {
	return stack.top == -1
}

// IsFull will return true if the stack is full and false otherwise
func (stack *StackUsingArray) IsFull() bool {
	return stack.top == stack.size-1
}

// StackTop will return true the top most element, panics if the stack is empty
func (stack *StackUsingArray) StackTop() int {
	if stack.IsEmpty() {
		panic(errors.New("Stack is empty"))
	}

	return stack.data[stack.top]
}

// String method will return a string representation of the stack
func (stack *StackUsingArray) String() string {

	if stack.IsEmpty() {
		return "{}"
	}

	var sb strings.Builder

	for i := 0; i < stack.top+1; i++ {
		sb.WriteString(strconv.Itoa(stack.data[stack.top-i]))
		sb.WriteByte('\n')
		if i != stack.top {
			sb.WriteRune('↓')
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}
