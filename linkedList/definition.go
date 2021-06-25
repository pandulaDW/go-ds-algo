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

// linkedList defines the structure of the linked list
type linkedList struct {
	head *node // points to the first element
	tail *node // points to the last element
	size int   // number of elements in the list
}

// NewLinkedList creates a new singly linked list
func NewLinkedList() *linkedList {
	head := &node{data: nil, pointer: nil}
	tail := &node{data: nil, pointer: nil}
	return &linkedList{head: head, tail: tail, size: 0}
}

// Size returns the total number of elements in the list
func (list *linkedList) Size() int {
	return list.size
}

// String implements the stringer interface
func (list *linkedList) String() string {
	strArr := make([]string, 0, list.size)
	currentNode := list.head.pointer

	for {
		strArr = append(strArr, fmt.Sprintf("%v", currentNode.data))
		currentNode = currentNode.pointer
		if currentNode == list.tail {
			break
		}
	}

	return fmt.Sprintf("%v\nLinkedList with %d elements", strings.Join(strArr, " -> "), list.size)
}
