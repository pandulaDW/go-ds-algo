package linkedlist

// Push elements at the end of the linked list
//
// Creates a new linked list struct and assigns data.
//
// Tail will be assigned to the new node and also the current tail's next will be assigned
// to the new node.
func (list *LinkedList) Push(data int) {
	newNode := Node{data: data, next: nil}

	// initial head case
	if (*list).head.next == nil {
		(*list).head.next = &newNode
		(*list).tail = &newNode
		(*list).count++
		return
	}

	// Set the next pointer of the current tail to the new node and set tail to the new node
	currentTailNode := (*list).tail
	(*currentTailNode).next = &newNode
	(*list).tail = &newNode
	(*list).count++
}
