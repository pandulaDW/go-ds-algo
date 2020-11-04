package linkedlist

// Push elements at the end of the linked list
//
// Creates a new linked list struct and assigns data.
//
// Tail will be assigned to the new node and the current tail's next will be assigned
// to the new node.
func (linkedList *LinkedList) Push(data int) {
	newNode := Node{data: data, next: nil}

	// initial head case
	if (*linkedList).head.next == nil {
		(*linkedList).head.next = &newNode
		(*linkedList).tail = &newNode
		return
	}

	// Set the next pointer of the tail to the new node and set tail to the new node
	currentTailNode := (*linkedList).tail
	(*currentTailNode).next = &newNode
	(*linkedList).tail = &newNode
}
