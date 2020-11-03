package linkedlist

// Push elements at the end of the linked list
//
// Creates a new linked list struct and assigns data.
// linked list will be iterated until a null pointer reference is found
// and the new node will be created and will be pointed to the last node
func (linkedList *LinkedList) Push(data int) {
	newNode := Node{data: data, next: nil}

	// initial head case
	if linkedList.head.next == nil {
		linkedList.head.next = &newNode
		return
	}

	// nodes after the head
	nextNode := linkedList.head
	for true {
		if nextNode.next == nil {
			nextNode.next = &newNode
			linkedList.tail = &newNode
			return
		}
		nextNode = nextNode.next
	}
}
