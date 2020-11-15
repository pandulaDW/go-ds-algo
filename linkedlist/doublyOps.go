package linkedlist

import (
	"strconv"
	"strings"
)

// contains few operations for doubly linked lists

// Push will push the elements to the doubly linked list
func (list *DoublyList) Push(data int) {
	newNode := DoublyNode{data: data, next: nil, previous: nil}

	// initial head case
	if list.head.next == nil {
		newNode.previous = list.head
		list.head.next = &newNode
		list.tail = &newNode
		list.count++
		return
	}

	// Set the next pointer of the current tail to the new node and set tail to the new node
	currentTailNode := list.tail
	newNode.previous = currentTailNode
	currentTailNode.next = &newNode
	list.tail = &newNode
	list.count++
}

// Pop will remove the last element from the doubly linked list in constant time
func (list *DoublyList) Pop() {
	newLastNode := list.tail.previous
	newLastNode.next = nil
	list.tail = newLastNode
	list.count--
}

// Returns a string representation of the linked list
// if reverse is set to true, the list will be displayed in a reverse fashion
func (list *DoublyList) String(reverse bool) string {
	// only head case
	if (*list).head.next == nil {
		return "{}"
	}

	// builder to build the linked list
	var sb strings.Builder
	var nextNode *DoublyNode

	if !reverse {
		nextNode = list.head
	} else {
		nextNode = list.tail
	}

	for true {
		if reverse {
			sb.WriteString(strconv.Itoa(nextNode.data))
			if nextNode == list.head.next {
				break
			} else {
				sb.WriteString(" <- ")
				nextNode = nextNode.previous
			}
		}

		if !reverse {
			nextNode = nextNode.next
			sb.WriteString(strconv.Itoa(nextNode.data))
			if nextNode.next == nil {
				break
			} else {
				sb.WriteString(" -> ")
			}
		}
	}

	return sb.String()
}
