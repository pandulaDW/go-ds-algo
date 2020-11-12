package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsParenthesisMatching(t *testing.T) {
	// correct case
	s1 := "((a-b) * (a-b))"
	assert.Equal(t, true, IsParenthesisMatching(s1))

	// More opening brackets
	s2 := "(((a-b) * (a-b))"
	assert.Equal(t, false, IsParenthesisMatching(s2))

	// More closing brackets
	s3 := "((a-b) * (a-b)))"
	assert.Equal(t, false, IsParenthesisMatching(s3))

	// multi type bracket correct
	s4 := "{([a+b] * [c-d])/e}"
	assert.Equal(t, true, IsParenthesisMatching(s4))

	// multi type bracket wrong
	s5 := "{([a+b) * [c-d]]/e}"
	assert.Equal(t, false, IsParenthesisMatching(s5))
}
