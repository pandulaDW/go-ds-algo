package trees

import (
	"fmt"

	"algos.com/main/queues"
)

// PreOrderRecursive will traverse a binary tree in a pre-order fashion recursively
func (root *Node) PreOrderRecursive() {
	var traverse func(t *Node)

	traverse = func(t *Node) {
		if t != nil {
			fmt.Printf("%d, ", t.Data)
			traverse(t.Left)
			traverse(t.Right)
		}
	}

	traverse(root)
}

// PreOrderIterative will traverse a binary tree in a pre-order fashion iteratively
func (root *Node) PreOrderIterative() {
	q := queues.CreateQueueUsingList()

	for {
		if q.IsEmpty() {
			break
		}
	}
}

// InOrderTraversal will traverse a binary tree in a pre-order fashion
// and print the data
func (root *Node) InOrderTraversal() {
	var traverse func(t *Node)

	traverse = func(t *Node) {
		if t != nil {
			traverse(t.Left)
			fmt.Printf("%d, ", t.Data)
			traverse(t.Right)
		}
	}

	traverse(root)
}
