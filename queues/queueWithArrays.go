package queues

import (
	"strconv"
	"strings"
)

// QueueUsingArray is the type of definition of a queue implemented using an array
type QueueUsingArray struct {
	data []int
}

// CreateQueueUsingArray creates a new queue using an array
func CreateQueueUsingArray() *QueueUsingArray {
	return new(QueueUsingArray)
}

// Size returns the number of elements of the queue
func (queue *QueueUsingArray) Size() int {
	return len(queue.data)
}

// Enqueue will insert a new element to the queue. Since slice appending is done
// at the end of the slice, same technique can be used for the queue
func (queue *QueueUsingArray) Enqueue(item int) {
	queue.data = append(queue.data, item)
}

// String implements the stringer interface for the queue
func (queue *QueueUsingArray) String() string {
	sb := make([]string, queue.Size())
	for i, val := range queue.data {
		sb[i] = strconv.Itoa(val)
	}
	return strings.Join(sb, " <- ")
}
