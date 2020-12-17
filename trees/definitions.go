package trees

// Node represents a node of the tree
type Node struct {
	left  *Node
	right *Node
	data  interface{}
}
