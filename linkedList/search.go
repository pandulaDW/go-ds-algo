package linkedList

// FindPredicate should take in a value and return a boolean
type FindPredicate func(val interface{}) bool

// FindIndex returns the index of first element found, by applying the given
// predicate function to each element's data.
//
// Returns -1, if no element has been found
func (list *linkedList) FindIndex(predicate FindPredicate) int {
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
