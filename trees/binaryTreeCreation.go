package trees

import "algos.com/main/queues"

// CreateTestTree will create a sample tree for a given set of data
// slice and it will return the root node
func CreateTestTree(data []int) *Node {
	q := queues.CreateQueueUsingList(len(data))
	root := Node{nil, nil, data[0]}
	q.Enqueue(&root)

	currentNode := q.Dequeue()
	return &root
}
