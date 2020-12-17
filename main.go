package main

import (
	"fmt"

	"algos.com/main/trees"
)

func main() {
	tree, err := trees.CreateTreeFromJSON("data/tree1.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(tree.Left.Data)
}
