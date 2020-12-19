package main

import (
	"algos.com/main/trees"
)

func main() {
	preOrder := []int{2, 1, 3, 5}
	InOrder := []int{1, 2, 5, 3, 6}
	trees.GenerateTreeFromTraversal(preOrder, InOrder)
}
