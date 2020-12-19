package trees

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

// CreateTreeFromJSON will create a binary tree from a JSON document
// and it will construct a tree and return the root node
func CreateTreeFromJSON(filepath string) (*Node, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	root := new(Node)

	err = json.Unmarshal(data, root)
	if err != nil {
		return nil, err
	}

	return root, nil
}

// GenerateTreeFromTraversal will take in a pre-order traversal list of nodes and an
// in-order traversal list of nodes and then will generate a unique tree based on that.
// Will return the root of that tree.
//
// Panics if the data is mismatching between the two traversals
func GenerateTreeFromTraversal(preOrder, inOrder []int) *Node {
	// Assert if elements are matching
	assert.ElementsMatch(nil, preOrder, inOrder)

	root := new(Node)
	return root
}
