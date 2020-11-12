package main

import "algos.com/main/stacks"

// IsParenthesisMatching checks if parenthesis are matching in a
// given expression.
//
// This function matche (), [] and {} types of parenthesis expressions
func IsParenthesisMatching(expr string) bool {

	oMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	cMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := stacks.CreateStackUsingArray(len(expr))
	var result bool

	for i, val := range expr {
		// if an opening bracket is found, push to the stack
		if _, ok := oMap[val]; ok {
			stack.Push(int(val))
		}
		// if a closing bracket is found, check it's corresponding opening bracket
		// if the stack is empty, return false
		if _, ok := cMap[val]; ok && stack.IsEmpty() {
			result = false
			break
		}
		// if the corresponding opening bracket is wrong, return false
		if _, ok := cMap[val]; ok && stack.StackTop() != int(cMap[val]) {
			result = false
			break
		}
		// if the corresponding opening bracket is correct, pop that bracket from stack
		if _, ok := cMap[val]; ok && stack.StackTop() == int(cMap[val]) {
			stack.Pop()
		}
		// at the end, if the stack is not empty, return false
		if i == len(expr)-1 && !stack.IsEmpty() {
			result = false
		}
		// at the end, if the stack is empty, return true
		if i == len(expr)-1 && stack.IsEmpty() {
			result = true
		}
	}

	return result
}
