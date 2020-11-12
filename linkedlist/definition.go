package linkedlist

// LinkedList is the type definition of a singly linked list
type LinkedList struct {
	head  *Node
	tail  *Node
	count int
}

// Node is the type definition of a node in a linked list
type Node struct {
	data int
	next *Node
}

// CreateLinkedList creates a new linked list with head and tail
// pointing to a common node
func CreateLinkedList() LinkedList {
	head := &Node{0, nil}
	tail := head

	list := LinkedList{head: head, tail: tail}
	return list
}

// IsEmpty would return if the list is empty
func (list *LinkedList) IsEmpty() bool {
	return list.head.next == nil
}
