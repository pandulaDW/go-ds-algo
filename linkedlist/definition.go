package linkedlist

// LinkedList is the type definition of a singly linked list
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// DoublyList is the type definition of a singly linked list
type DoublyList struct {
	head  *DoublyNode
	tail  *DoublyNode
	count int
}

// Node is the type definition of a node in a linked list
type Node struct {
	data int
	next *Node
}

// DoublyNode is the type definition of a node in a doubly linked list
type DoublyNode struct {
	data     int
	next     *DoublyNode
	previous *DoublyNode
}

// CreateLinkedList creates a new linked list with head and tail
// pointing to a common node
func CreateLinkedList() LinkedList {
	head := &Node{0, nil}
	tail := head

	list := LinkedList{head: head, tail: tail}
	return list
}

// CreateDoublyLinkedList creates a new doubly linked list with head and tail
// pointing to a common node
func CreateDoublyLinkedList() DoublyList {
	head := &DoublyNode{data: 0, next: nil, previous: nil}
	tail := head
	return DoublyList{head: head, tail: tail, count: 0}
}

// IsEmpty would return if the list is empty
func (list *LinkedList) IsEmpty() bool {
	return list.head.next == nil
}
