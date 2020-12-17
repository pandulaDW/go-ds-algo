package queues

import (
	"errors"
	"strings"
)

// QueueUsingList is the type of definition of a queue implemented using a linked list
type QueueUsingList struct {
	data *LinkedList
}

// CreateQueueUsingList creates a new queue using a linked list
func CreateQueueUsingList() *QueueUsingList {
	queue := &QueueUsingList{}
	queue.data = CreateLinkedList()
	return queue
}

// Size returns the number of elements of the queue
func (queue *QueueUsingList) Size() int {
	return queue.data.Size()
}

// IsEmpty returns whether the queue is empty or not
func (queue *QueueUsingList) IsEmpty() bool {
	return queue.Size() == 0
}

// IsFull returns whether the queue is at capacity
func (queue *QueueUsingList) IsFull() bool {
	return false
}

// First returns the first item of the queue
func (queue *QueueUsingList) First() interface{} {
	return queue.data.First()
}

// Last returns the last item of the queue
func (queue *QueueUsingList) Last() interface{} {
	return queue.data.Last()
}

// Enqueue will insert a new element to the queue. Push method of the linked list
// is used here. Panics if exceeding the queue capacity
func (queue *QueueUsingList) Enqueue(item interface{}) {
	if queue.IsFull() {
		panic(errors.New("Cannot enqueue to a full queue"))
	}
	queue.data.Push(item)
}

// Dequeue will remove an element to the queue. Pull method of the linked list
// is used here. Panics if queue is empty
func (queue *QueueUsingList) Dequeue() interface{} {
	if queue.IsEmpty() {
		panic(errors.New("Cannot dequeue an empty queue"))
	}
	first := queue.First()
	queue.data.Pull()
	return first
}

// String implements the stringer interface for the queue
func (queue *QueueUsingList) String() string {
	list := queue.data.String()
	listData := strings.Split(list, " -> ")
	return strings.Join(listData, " <- ")
}
