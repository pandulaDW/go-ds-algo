package queues

import (
	"errors"
	"strings"
)

// QueueUsingArray is the type of definition of a queue implemented using an array
type QueueUsingArray struct {
	data     []interface{}
	capacity int
}

// CreateQueueUsingArray creates a new queue using an array
func CreateQueueUsingArray(capacity int) *QueueUsingArray {
	queue := &QueueUsingArray{}
	queue.data = make([]interface{}, 0, capacity)
	queue.capacity = capacity
	return queue
}

// Size returns the number of elements of the queue
func (queue *QueueUsingArray) Size() int {
	return len(queue.data)
}

// IsEmpty returns whether the queue is empty or not
func (queue *QueueUsingArray) IsEmpty() bool {
	return queue.Size() == 0
}

// IsFull returns whether the queue is at capacity
func (queue *QueueUsingArray) IsFull() bool {
	return queue.Size() == queue.capacity
}

// First returns the first item of the queue
func (queue *QueueUsingArray) First() interface{} {
	return queue.data[0]
}

// Last returns the last item of the queue
func (queue *QueueUsingArray) Last() interface{} {
	return queue.data[queue.Size()-1]
}

// Enqueue will insert a new element to the queue. Since slice appending is done
// at the end of the slice, same technique can be used for the queue. Panics if
// exceeding the queue capacity
func (queue *QueueUsingArray) Enqueue(item interface{}) {
	if queue.IsFull() {
		panic(errors.New("Cannot enqueue to a full queue"))
	}
	queue.data = append(queue.data, item)
}

// Dequeue will remove an element to the queue. Changing the underlying slice is
// sufficient here. Will be an O(1) operation. Panics if queue is empty
func (queue *QueueUsingArray) Dequeue() {
	if queue.IsEmpty() {
		panic(errors.New("Cannot dequeue an empty queue"))
	}
	queue.data = queue.data[1:]
}

// String implements the stringer interface for the queue
func (queue *QueueUsingArray) String() string {
	sb := make([]string, queue.Size())
	for i, val := range queue.data {
		sb[i] = val.(string)
	}
	return strings.Join(sb, " <- ")
}
