package linkedlist

import (
	"strconv"
	"strings"
)

// Returns a string representation of the linked list
func (linkedList *LinkedList) String() string {

	// only head case
	if linkedList.head.next == nil {
		return "{}"
	}

	// builder to build the linked list
	var sb strings.Builder

	nextNode := linkedList.head
	for true {
		sb.WriteString(strconv.FormatInt(int64(nextNode.data), 10))
		if nextNode.next == nil {
			sb.WriteByte('\n')
			break
		} else {
			sb.WriteString(" -> ")
		}
		nextNode = nextNode.next
	}

	return sb.String()
}
