package trees

// Node represents a node of the tree
type Node struct {
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
	Data  int   `json:"data"`
}
