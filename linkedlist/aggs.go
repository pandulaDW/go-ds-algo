package linkedlist

import "math"

// Count returns the count of the linked list struct
func (list *LinkedList) Count() int {
	return (*list).count
}

// Sum returns the sum of the linked list
//
// Returns 0 if only head element is there
func (list *LinkedList) Sum() int {

	// only head case
	if list.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*list).head.next
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
func (list *LinkedList) Max() int {

	// only head case
	if list.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*list).head.next
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
func (list *LinkedList) Min() int {

	// only head case
	if list.head.next == nil {
		return 0
	}

	// iterate the list
	nextNode := (*list).head.next
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
