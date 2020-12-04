package main

import (
	"fmt"

	"algos.com/main/queues"
)

func main() {
	q := queues.CreateQueueUsingArray(10)
	q.Enqueue(12)
	q.Enqueue(15)
	q.Enqueue(19)
	q.Enqueue(22)
	q.Enqueue(49)

	q.Dequeue()
	q.Dequeue()

	fmt.Println("First element: ", q.First())
	fmt.Println("Last element: ", q.Last())

	fmt.Println(q.String())
}
