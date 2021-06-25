package linkedList

// FindPredicate should take in a value and return a boolean
type FindPredicate func(val interface{}) bool

// FindIndex returns the index of first element found, by applying the given
// predicate function to each element's data.
//
// Returns -1, if no element has been found
func (list *LinkedList) FindIndex(predicate FindPredicate) int {
	currentNode := list.head.pointer
	if currentNode == nil {
		return -1
	}

	index := 0
	for {
		if currentNode == list.tail {
			break
		}
		if predicate(currentNode.data) {
			return index
		}
		currentNode = currentNode.pointer
		index++
	}

	return -1
}

// Find returns the first element found, by applying the given
// predicate function to each element's data.
//
// Returns nil, if no element has been found
func (list *LinkedList) Find(predicate FindPredicate) interface{} {
	currentNode := list.head.pointer
	if currentNode == nil {
		return nil
	}

	index := 0
	for {
		if currentNode == list.tail {
			break
		}
		if predicate(currentNode.data) {
			return currentNode.data
		}
		currentNode = currentNode.pointer
		index++
	}

	return -1
}
