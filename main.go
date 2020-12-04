package main

import (
	"fmt"

	"algos.com/main/queues"
)

func main() {
	q := queues.CreateQueueUsingArray()
	q.Enqueue(12)
	q.Enqueue(15)
	q.Enqueue(19)
	q.Enqueue(22)

	fmt.Println(q.String())
}
