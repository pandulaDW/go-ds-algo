package linkedlist

import "errors"

// Pop will remove the last element from the list adjust the tail position
//
// Returns an error if the head is the last element and pop is called
func (linkedList *LinkedList) Pop() error {

	// handle only head case
	if (*linkedList).head == (*linkedList).tail {
		return errors.New("Cannot remove head from the ")
	}

	// point the tail to the previous

	return nil
}
