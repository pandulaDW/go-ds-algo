package main

import (
	"fmt"

	"algos.com/main/trees"
)

func main() {
	root, err := trees.CreateTreeFromJSON("data/tree2.json")
	if err != nil {
		panic(err)
	}

	root.PreOrderRecursive()
	fmt.Println()

	root.PreOrderIterative()
	fmt.Println()

	root.InOrderRecursive()
	fmt.Println()

	root.InOrderIterative()
	fmt.Println()

	root.LevelOrderIterative()
}
