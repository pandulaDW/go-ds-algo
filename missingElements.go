package main

import (
	"fmt"

	"algos.com/main/helpers"
)

// finding missing elements in a sorted array of natural numbers
// O(n) solution

func findSingleMissingElements(arr []int) []int {
	missingElements := make([]int, 0, cap(arr))
	index := 0
	correctDiff := arr[0] - index

	for i := 0; i < len(arr); i++ {
		difference := arr[i] - index
		if difference != correctDiff {
			missingElements = append(missingElements, arr[i]-1)
			correctDiff = difference
		}
		index++
	}
	return missingElements
}

// takes O(n2) time
func findMultipleMissingElements(arr []int) []int {
	missingElements := make([]int, 0, cap(arr))
	index := 0
	correctDiff := arr[0] - index

	for i := 0; i < len(arr); i++ {
		difference := arr[i] - index
		offset := difference - correctDiff
		if offset > 0 {
			if offset > 1 {
				for j := 1; j < offset+1; j++ {
					missingElements = append(missingElements, arr[i]-j)
					correctDiff = difference
				}
			} else {
				missingElements = append(missingElements, arr[i]-1)
				correctDiff = difference
			}
		}
		index++
	}
	return missingElements
}

// find missing elements in an unsorted array using another list
// takes O(3n) time
func findMissingElementsUnsorted(arr []int) []int {
	maxEl := helpers.GetMaxElem(arr) // order n
	hashArray := make([]int, maxEl)

	for _, val := range arr { // order n
		hashArray[val-1] = 1
	}

	fmt.Println(hashArray)

	missingElems := make([]int, 0)
	for i, val := range hashArray { // order n
		if val == 0 {
			missingElems = append(missingElems, i)
		}
	}

	return missingElems
}
