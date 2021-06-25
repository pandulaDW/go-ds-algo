package linkedList

// InsertAtEnd appends the given data to the end of the linked list
func (list *LinkedList) InsertAtEnd(data interface{}) {
	newNode := &node{data: data, pointer: list.tail}
	list.size++

	// if initial element
	if list.head.pointer == nil {
		list.head.pointer = newNode
		list.tail.pointer = newNode
		return
	}

	currLastNode := list.tail.pointer
	currLastNode.pointer = newNode
	list.tail.pointer = newNode
}

// InsertAtStart prepends the given data to the start of the linked list
func (list *LinkedList) InsertAtStart(data interface{}) {
	newNode := &node{data: data}
	list.size++

	// if initial element
	if list.head.pointer == nil {
		list.head.pointer = newNode
		list.tail.pointer = newNode
		return
	}

	currFirstNode := list.head.pointer
	newNode.pointer = currFirstNode
	list.head.pointer = newNode
}

// InsertAtPosition adds the element at the specified index and shifts the previously occupied
// element to the right.
//
// If the index is higher than the size of the list, it will appended to the end.
func (list *LinkedList) InsertAtPosition(index uint, data interface{}) {
	i := int(index)
	newNode := &node{data: data}
	list.size++

	if i >= list.size-1 {
		list.InsertAtEnd(data)
		return
	}

	if i == 0 {
		list.InsertAtStart(data)
		return
	}

	currNode := list.head.pointer
	currIndex := 1

	for {
		if currIndex == i {
			newNode.pointer = currNode.pointer
			currNode.pointer = newNode
			break
		}
		currIndex++
		currNode = currNode.pointer
	}
}
