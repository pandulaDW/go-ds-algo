package linkedlist

// ReverseList reverses a linked list in an O(n) time and space complexity
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

	// iterate the array and replace the data values in the list
	nextNode = list.head.next
	for _, val := range reversedArr {
		nextNode.data = val
		nextNode = nextNode.next
	}
}
