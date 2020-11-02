package main

import "strings"

// Returns the length of the last word of a
// sentence. The below solution is an O(1) solution
func getLastWordCount(phrase string) int {
	phrase = strings.TrimSpace(phrase)
	count := 0
	for i := len(phrase) - 1; i > -1; i-- {
		if string(phrase[i]) != " " {
			count++
		} else {
			break
		}
	}
	return count
}
