package linkedlist

// DummyLoop creates a loop from a linked list
func DummyLoop() LinkedList {
	list := CreateLinkedList()
	list.Push(10)
	list.Push(20)
	list.Push(30)
	list.Push(40)
	list.Push(50)

	// make tail point to the first node
	list.tail.next = list.head.next
	return list
}

// IsLoop checks whether the linked list is a loop
//
// Returns true if it's a loop and false if not
func (list *LinkedList) IsLoop() bool {

	// empty case
	if list.head.next == nil {
		return false
	}

	// create a map to hold addresses of nodes
	elemMap := map[*Node]int{}

	// iterate until the list is exhausted or an address is found again
	nextNode := list.head.next

	for nextNode != nil {
		if _, ok := elemMap[nextNode]; !ok {
			elemMap[nextNode] = 0
		} else {
			return true
		}
		nextNode = nextNode.next
	}

	return false
}
