package linkedlist

// ReverseList reverses a linked list in an O(2n) complexity
func (list *LinkedList) ReverseList() {

	// handle empty list
	if list.head.next == nil {
		return
	}

	// iterate the list and save the results in an array
	reversedArr := make([]int, list.count)
	nextNode := list.head.next
	index := 1

	for index <= list.count {
		reversedArr[list.count-index] = nextNode.data
		index++
		nextNode = nextNode.next
	}

	// create a new list
	head := &Node{0, nil}
	tail := head
	reversedList := LinkedList{head: head, tail: tail}

	// push the reversed array elements to the new linked list
	for _, val := range reversedArr {
		reversedList.Push(val)
	}

	// assign the new linked list to the old list
	*list = reversedList
}
