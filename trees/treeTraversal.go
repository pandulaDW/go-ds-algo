package trees

import "fmt"

// PreOrderTraversal will traverse a binary tree in a pre-order fashion
// and print the data
func (root *Node) PreOrderTraversal() {
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
