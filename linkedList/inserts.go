package linkedList

// InsertAtEnd appends the given data to the end of the linked list
func (list *linkedList) InsertAtEnd(data interface{}) {
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
