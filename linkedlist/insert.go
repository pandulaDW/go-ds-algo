package linkedlist

import "errors"

// Push elements to the end of the linked list
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

// InsertFirst inserts given data at the beginning of the linked list
func (list *LinkedList) InsertFirst(data int) {

	// if the length is 0, use push method
	if (*list).head.next == nil {
		list.Push(data)
		return
	}

	// Create a new node and point it to the current first element
	currentFirstElement := (*list).head.next
	newNode := Node{data: data, next: currentFirstElement}

	// Assigns the list's head to the new node
	(*list).head.next = &newNode

	// increase count
	(*list).count++
}

// InsertAtIndex inserts given data at the given index of the linked list
//
// returns error if the index is equal or higher than list size
//
// To insert at the end of the list, use the Push method
func (list *LinkedList) InsertAtIndex(index, data int) error {

	// if the index is 0, use InsertFirst method
	if index == 0 {
		list.InsertFirst(data)
		return nil
	}
	// if index is higher, return indexOutOfBound error
	if index >= list.count || index < 0 {
		return errors.New("Index out of Bound")
	}

	// create a new node
	newNode := Node{data: data, next: nil}

	// find currentNode and previousToCurrentNode for the given index
	var currentNode, previousToCurrentNode *Node
	currentNode = (*list).head.next
	counter := 0

	for true {
		if counter == index {
			newNode.next = currentNode
			(*previousToCurrentNode).next = &newNode
			(*list).count++
			break
		}

		if counter == index-1 {
			previousToCurrentNode = currentNode
		}

		counter++
		currentNode = (*currentNode).next
	}

	return nil
}

// PushToSorted pushes elements to the linked list while preserving the sort order
//
// The method assumes that the list is sorted in an ascending order and only unique elements
// are allowed
func (list *LinkedList) PushToSorted(data int) {

	// if the list is empty
	if list.count == 0 {
		list.Push(data)
	}

	// find the closest index position for the new node
	nextNode := list.head.next
	correctIndex := 0

	for correctIndex < (*list).count {
		if (*nextNode).data == data {
			return
		}

		if (*nextNode).data > data {
			list.InsertAtIndex(correctIndex, data)
			return
		}

		correctIndex++
		nextNode = (*nextNode).next
	}

	// if the element is higher than current elements, push it to the end
	list.Push(data)
}
