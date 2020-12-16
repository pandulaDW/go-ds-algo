package queues

// Queue interface defines the methods that should be included in a
// queue implementation
type Queue interface {
	Size() int
	IsEmpty() bool
	IsFull() bool
	First() int
	Last() int
	Enqueue(item int)
	Dequeue()
	String() string
}
