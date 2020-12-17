package queues

// Queue interface defines the methods that should be included in a
// queue implementation
type Queue interface {
	Size() int
	IsEmpty() bool
	IsFull() bool
	First() interface{}
	Last() interface{}
	Enqueue(item interface{})
	Dequeue()
	String() string
}
