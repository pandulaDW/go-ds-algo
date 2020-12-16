package queues

import (
	"fmt"
	"strings"
)

// A singly linked list defnition with minimal a number of methods
// to implement a queue. Only pusing at the end and pulling from the
// front is implemented

// Node defines a generic data node in the list
type Node struct {
	data int
	next *Node
}

// HeadNode defines the head of the list
type HeadNode struct {
	next *Node
}

// TailNode defines the head of the list
type TailNode struct {
	previous *Node
}

// LinkedList struct
type LinkedList struct {
	head  *HeadNode
	tail  *TailNode
	count int
}

// CreateLinkedList will create a new linked list
func CreateLinkedList() *LinkedList {
	return &LinkedList{&HeadNode{}, &TailNode{}, 0}
}

// Push method will push an element to the list at the end in constant time
func (list *LinkedList) Push(item int) {

	// define a new node
	newNode := &Node{data: item, next: nil}

	if list.head.next == nil {
		// push to head if it's the first element
		list.head.next = newNode
	} else {
		// set the current last element's next as the newNode
		list.tail.previous.next = newNode
	}

	// update the tail
	list.tail.previous = newNode
	list.count++
}

// String implements the stringer interface
func (list *LinkedList) String() string {
	sb := make([]string, 0, list.count)

	if list.head.next == nil {
		return "{}"
	}

	currentNode := list.head.next
	for {
		sb = append(sb, fmt.Sprintf("%d", currentNode.data))
		currentNode = currentNode.next
		if currentNode == nil {
			break
		}
	}

	return strings.Join(sb, " -> ")
}
