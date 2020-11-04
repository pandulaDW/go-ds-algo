package linkedlist

import (
	"strconv"
	"strings"
)

// Returns a string representation of the linked list
func (list *LinkedList) String() string {

	// only head case
	if (*list).head.next == nil {
		return "{}"
	}

	// builder to build the linked list
	var sb strings.Builder

	nextNode := (*list).head
	for true {
		nextNode = nextNode.next
		sb.WriteString(strconv.FormatInt(int64(nextNode.data), 10))
		if nextNode.next == nil {
			break
		} else {
			sb.WriteString(" -> ")
		}
	}

	return sb.String()
}
