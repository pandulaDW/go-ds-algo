package trees

import (
	"fmt"
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
