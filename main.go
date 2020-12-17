package main

import (
	"fmt"

	"algos.com/main/trees"
)

func main() {
	root, err := trees.CreateTreeFromJSON("data/tree1.json")
	if err != nil {
		panic(err)
	}

	fmt.Println()
	root.PreOrderRecursive()
}
