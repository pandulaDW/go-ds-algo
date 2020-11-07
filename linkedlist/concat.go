package linkedlist

// Concat will append the given list to the current list
func (list *LinkedList) Concat(list2 *LinkedList) {
	list.tail.next = list2.head.next
	list.count = list.count + list2.count
	list.tail = list2.tail
}

// Merge will merge the current sorted list with another sorted list
func (list *LinkedList) Merge(list2 *LinkedList) {

	// handle empty list cases
	if list.head.next == nil || list2.head.next == nil {
		list.Concat(list2)
	}

	// have two pointers to track the two lists.
	// p for longer list and q for shorter list
	p, q := list.head.next, list2.head.next

	if list.count < list2.count {
		p, q = list2.head.next, list.head.next
	}

	// create a new list
	mergedList := CreateLinkedList()

	// iterate until short list is exhausted
	for q != nil {
		if p != nil && p.data < q.data {
			mergedList.Push(p.data)
			p = p.next
		} else {
			mergedList.Push(q.data)
			q = q.next
		}
	}

	// if the long list has more elements to add, add them to the end
	for p != nil {
		mergedList.Push(p.data)
		p = p.next
	}

	// assign the current list to the merged list
	*list = mergedList
}
