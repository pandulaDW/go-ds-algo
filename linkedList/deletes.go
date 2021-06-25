package linkedList

// RemoveFirst will remove the first node from the linked list
func (list *LinkedList) RemoveFirst() {
	if list.size == 0 {
		return
	}
	list.size--

	currentFirstNode := list.head.pointer
	newFirstNode := currentFirstNode.pointer

	currentFirstNode.pointer = nil
	list.head.pointer = newFirstNode
}

// RemoveLast will remove the last node from the linked list.
//
// Since the linked list is singly, deletion from tail takes O(n) time
func (list *LinkedList) RemoveLast() {
	if list.size == 0 {
		return
	}

	if list.size == 1 {
		list.head.pointer.pointer = nil
		list.head.pointer = nil
		list.tail.pointer = nil
		list.size--
		return
	}

	list.size--
	currentNode := list.head.pointer
	for {
		if currentNode.pointer.pointer == list.tail {
			currentNode.pointer.pointer = nil
			currentNode.pointer = list.tail
			list.tail.pointer = currentNode.pointer
			break
		}
		currentNode = currentNode.pointer
	}
}

// RemoveUsingIndex will remove the element in the given index.
func (list *LinkedList) RemoveUsingIndex(index uint) {
	i := int(index)

	if i > list.size-1 {
		return
	}

}
