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
	s := CreateSimpleStack()
	currentNode := root

	for {
		if currentNode == nil && s.IsEmpty() {
			break
		}

		if currentNode != nil {
			fmt.Printf("%d, ", currentNode.Data)
			s.Push(currentNode)
			currentNode = currentNode.Left
		} else {
			currentNode = s.Pop().(*Node)
			currentNode = currentNode.Right
		}
	}
}

// InOrderRecursive will traverse a binary tree in an in-order fashion recursively
func (root *Node) InOrderRecursive() {
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

// InOrderIterative will traverse a binary tree in an in-order fashion iteratively
func (root *Node) InOrderIterative() {
	s := CreateSimpleStack()
	currentNode := root

	for {
		if currentNode == nil && s.IsEmpty() {
			break
		}

		if currentNode != nil {
			s.Push(currentNode)
			currentNode = currentNode.Left
		} else {
			currentNode = s.Pop().(*Node)
			fmt.Printf("%d, ", currentNode.Data)
			currentNode = currentNode.Right
		}
	}
}

// LevelOrderIterative will traverse a binary tree in a level-order fashion iteratively
func (root *Node) LevelOrderIterative() {
	q := queues.CreateQueueUsingList()
	fmt.Printf("%d, ", root.Data)
	q.Enqueue(root)
	currentNode := root

	for !q.IsEmpty() {
		currentNode = q.Dequeue().(*Node)

		if currentNode.Left != nil {
			fmt.Printf("%d, ", currentNode.Left.Data)
			q.Enqueue(currentNode.Left)
		}

		if currentNode.Right != nil {
			fmt.Printf("%d, ", currentNode.Right.Data)
			q.Enqueue(currentNode.Right)
		}
	}
}
