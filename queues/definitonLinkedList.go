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
	data interface{}
	next *Node
}

// HeadNode defines the head of the list
type HeadNode struct {
	first *Node
}

// TailNode defines the head of the list
type TailNode struct {
	last *Node
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

// Size is a getter for list count property
func (list *LinkedList) Size() int {
	return list.count
}

// First returns the first element of the list
func (list *LinkedList) First() interface{} {
	return list.head.first.data
}

// Last returns the last element of the list
func (list *LinkedList) Last() interface{} {
	return list.tail.last.data
}

// Push method will push an element to the list at the end in constant time
func (list *LinkedList) Push(item interface{}) {

	// define a new node
	newNode := &Node{data: item, next: nil}

	if list.head.first == nil {
		// push to head if it's the first element
		list.head.first = newNode
	} else {
		// set the current last element's next as the newNode
		list.tail.last.next = newNode
	}

	// update the tail
	list.tail.last = newNode
	list.count++
}

// Pull method will pull the first element from the list in constant time.
func (list *LinkedList) Pull() {
	if list.head.first == nil {
		return
	}

	// set the head correctly
	currentHead := list.head.first
	list.head.first = currentHead.next

	// set the current head's pointer to null for it to be gced
	currentHead.next = nil

	// set the list count
	list.count--
}

// String implements the stringer interface
func (list *LinkedList) String() string {
	sb := make([]string, 0, list.count)
	currentNode := list.head.first

	if currentNode == nil {
		return "{}"
	}

	for {
		sb = append(sb, fmt.Sprintf("%d", currentNode.data))
		currentNode = currentNode.next
		if currentNode == nil {
			break
		}
	}

	return strings.Join(sb, " -> ")
}
