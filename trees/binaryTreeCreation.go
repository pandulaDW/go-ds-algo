package trees

import (
	"encoding/json"
	"io/ioutil"
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
