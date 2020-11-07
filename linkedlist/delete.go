package linkedlist

import (
	"errors"
)

// Pop will remove the last element from the list adjust the tail position
//
// Returns an error if the head is the last element and pop is called
func (list *LinkedList) Pop() error {

	// handle only head case
	if (*list).head == (*list).tail {
		return errors.New("Cannot remove head from the ")
	}

	// point the tail to the previous

	return nil
}

// DeleteFirst will remove the element first element from the list
//
// Returns an error if the list is empty
func (list *LinkedList) DeleteFirst() error {

	if list.count == 0 {
		return errors.New("Index out of bound error")
	}

	firstElement := list.head.next
	list.head.next = firstElement.next

	list.count--

	return nil
}

// DeleteAtIndex will remove the element at the given index
//
// Returns an error if the index is higher than list size
func (list *LinkedList) DeleteAtIndex(index int) error {

	// if index is higher, return indexOutOfBound error
	if index >= list.count || index < 0 {
		return errors.New("Index out of Bound")
	}

	if index == 0 {
		list.DeleteFirst()
		return nil
	}

	var correctNode, previousToCorrectNode *Node
	correctNode = (*list).head.next
	counter := 0

	for true {
		if index == counter {
			(*previousToCorrectNode).next = (*correctNode).next
			(*list).count--
			break
		}

		if counter == index-1 {
			previousToCorrectNode = correctNode
		}

		counter++
		correctNode = correctNode.next
	}

	return nil
}

// DeleteDuplicates deletes duplicate entries from a linked list with O(n) complexity
//
// Returns an error if array is empty
func (list *LinkedList) DeleteDuplicates() error {

	if list.head.next == nil {
		return errors.New("Array cannot be empty")
	}

	// create an element map
	elemMap := map[int]int{}

	// iterate the list and if duplicates are found remove them
	index := 0
	nextNode := list.head.next
	previousNode := list.head

	for nextNode != nil {
		if _, ok := elemMap[nextNode.data]; !ok {
			elemMap[nextNode.data] = 0
			previousNode = previousNode.next
		} else {
			previousNode.next = nextNode.next
			nextNode.next = nil
			nextNode = previousNode
		}
		index++
		nextNode = nextNode.next
	}

	return nil
}
