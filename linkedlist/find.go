package linkedlist

import "errors"

// FindByIndex returns the data of a node at the given index
// position of the linked list
//
// if the index is higher than the list size, it will return -1 with an error
func (list *LinkedList) FindByIndex(index int) (int, error) {

	if list.count < index+1 {
		return -1, errors.New("Index cannot be larger than array size")
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
