package linkedlist

import (
	"errors"
)

// FindByIndex returns the data of a node at the given index
// position of the linked list
//
// if the index is higher than the list size, it will return -1 with an error
func (list *LinkedList) FindByIndex(index int) (int, error) {

	if list.count < index+1 {
		return -1, errors.New("Index out of bound")
	}

	nextNode := (*list).head.next
	data, currentIdx := 0, 0

	for true {
		if index == currentIdx {
			data = (*nextNode).data
			break
		}
		currentIdx++
		nextNode = (*nextNode).next
	}

	return data, nil
}

// FindByData returns the index of the first node found if the passed data
// matches the node data property.
//
// returns -1 if a matching node is not found
func (list *LinkedList) FindByData(data int) int {

	if (*list).head.next == nil {
		return -1
	}

	nextNode := (*list).head.next
	index := 0

	for index < list.count {
		if data == (*nextNode).data {
			return index
		}
		index++
		nextNode = (*nextNode).next
	}

	return -1
}

// FindByDataRepeated returns the index of the first node found if the passed data
// matches the node data property.
//
// The difference with FindByData is that, this procedure will move the matching node to the
// beginning of the list, thereby improving searching speed on repeated searches
//
// returns -1 if a matching node is not found
func (list *LinkedList) FindByDataRepeated(data int) int {
	if (*list).head.next == nil {
		return -1
	}

	nextNode := (*list).head.next
	previousNode := (*list).head
	index := 0

	for index < list.count {
		if data == (*nextNode).data {
			(*previousNode).next = (*nextNode).next
			(*nextNode).next = (*list).head.next
			newHead := Node{0, nextNode}
			(*list).head = &newHead

			// if the last element is matched, tail should be updated
			if index == list.count-1 {
				(*list).tail = previousNode
			}

			return index
		}
		index++
		previousNode = nextNode
		nextNode = (*nextNode).next
	}

	return -1
}

// FindLast will return the last element of the linked list, panics if list is empty
func (list *LinkedList) FindLast() int {
	if list.head.next == nil {
		panic(errors.New("List is empty"))
	}
	return list.tail.data
}

// FindFirst will return the first element of the linked list, panics if list is empty
func (list *LinkedList) FindFirst() int {
	if list.head.next == nil {
		panic(errors.New("List is empty"))
	}
	return list.head.next.data
}
