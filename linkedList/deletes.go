package linkedList

// removeOnlyRemainingElement removes the only remaining element from the list
func (list *LinkedList) removeOnlyRemainingElement() {
	list.head.pointer.pointer = nil
	list.head.pointer = nil
	list.tail.pointer = nil
}

// RemoveFirst will remove the first node from the linked list
func (list *LinkedList) RemoveFirst() {
	if list.size == 0 {
		return
	}
	list.size--

	if list.size == 0 {
		list.removeOnlyRemainingElement()
		return
	}

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
	list.size--

	if list.size == 0 {
		list.removeOnlyRemainingElement()
		return
	}

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

	if i == 0 {
		list.RemoveFirst()
		return
	}

	if i >= list.size-1 {
		list.RemoveLast()
		return
	}

	currentNode := list.head.pointer
	currentIndex := 1

	for {
		if currentIndex == i {
			removedNode := currentNode.pointer
			currentNode.pointer = removedNode.pointer
			removedNode.pointer = nil
			break
		}
		currentNode = currentNode.pointer
		currentIndex++
	}

	list.size--
}
