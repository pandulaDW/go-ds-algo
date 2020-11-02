package main

import (
	"testing"

	"algos.com/main/helpers"
	"github.com/stretchr/testify/assert"
)

func TestLastWordCount(t *testing.T) {
	s1 := "Hello World"
	s2 := "   word   "
	s3 := ""

	assert.Equal(t, 5, getLastWordCount(s1))
	assert.Equal(t, 4, getLastWordCount(s2))
	assert.Equal(t, 0, getLastWordCount(s3))
}

func TestFindSingleMissingElement(t *testing.T) {
	a := []int{6, 7, 8, 9, 10, 12, 13, 14, 15, 17, 18}
	assert.ElementsMatch(t, []int{11, 16}, findSingleMissingElements(a))
}

func TestFindMultipleMissingElement(t *testing.T) {
	a := []int{6, 7, 8, 9, 11, 12, 15, 16, 17, 18, 19}
	assert.ElementsMatch(t, []int{10, 13, 14}, findMultipleMissingElements(a))
}

func TestFindMissingElementsUnsorted(t *testing.T) {
	a := []int{7, 5, 10, 12, 3, 2, 16, 4}

	// test helper function max
	assert.Equal(t, 16, helpers.GetMaxElem(a))

	// assert.ElementsMatch(t, []int{1, 6, 8, 9, 11, 13, 14, 15}, findMissingElementsUnsorted(a))
}

func TestFindDuplicateElements(t *testing.T) {
	a := []int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20}
	assert.ElementsMatch(t, []int{8, 15}, findDuplicateElements(a))
}
