package linkedlist

import "math"

// Count returns the count of the linked list struct
func (linkedList *LinkedList) Count() int {
	return (*linkedList).count
}

// Sum returns the sum of the linked list
//
// Returns 0 if only head element is there
func (linkedList *LinkedList) Sum() int {

	// only head case
	if linkedList.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*linkedList).head.next
	sum := 0
	for true {
		sum += (*nextNode).data
		if (*nextNode).next == nil {
			break
		}
		nextNode = (*nextNode).next
	}
	return sum
}

// Max returns the max of the linked list
//
// Returns 0 if only head element is there
func (linkedList *LinkedList) Max() int {

	// only head case
	if linkedList.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*linkedList).head.next
	max := math.MinInt32
	for true {
		if max < (*nextNode).data {
			max = (*nextNode).data
		}
		if (*nextNode).next == nil {
			break
		}
		nextNode = (*nextNode).next
	}
	return max
}

// Min returns the min of the linked list
//
// Returns 0 if only head element is there
func (linkedList *LinkedList) Min() int {

	// only head case
	if linkedList.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*linkedList).head.next
	min := math.MaxInt32
	for true {
		if min > (*nextNode).data {
			min = (*nextNode).data
		}
		if (*nextNode).next == nil {
			break
		}
		nextNode = (*nextNode).next
	}
	return min
}
