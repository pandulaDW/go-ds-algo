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
		if currentNode != nil {
			fmt.Printf("%d, ", currentNode.Data)
			s.Push(currentNode)
			currentNode = currentNode.Left
			continue
		}

		currentNode = s.Pop().(*Node)

		if currentNode.Right == nil {
			if !s.IsEmpty() {
				currentNode = s.Pop().(*Node)
			} else {
				break
			}
			currentNode = currentNode.Right
			continue
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
