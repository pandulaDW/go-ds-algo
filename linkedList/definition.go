package linkedList

import (
	"fmt"
	"strings"
)

// node defines an internal node of the list
type node struct {
	data    interface{}
	pointer *node
}

// LinkedList defines the structure of the linked list
type LinkedList struct {
	head *node // points to the first element
	tail *node // points to the last element
	size int   // number of elements in the list
}

// NewLinkedList creates a new singly linked list
func NewLinkedList() *LinkedList {
	head := &node{data: nil, pointer: nil}
	tail := &node{data: nil, pointer: nil}
	return &LinkedList{head: head, tail: tail, size: 0}
}

// Size returns the total number of elements in the list
func (list *LinkedList) Size() int {
	return list.size
}

// String implements the stringer interface
func (list *LinkedList) String() string {
	if list.size == 0 {
		return "[]"
	}

	strArr := make([]string, 0, list.size)
	currentNode := list.head.pointer

	for {
		strArr = append(strArr, fmt.Sprintf("%v", currentNode.data))
		currentNode = currentNode.pointer
		if currentNode == list.tail {
			break
		}
	}

	return "[ " + strings.Join(strArr, " -> ") + " ]"
}
