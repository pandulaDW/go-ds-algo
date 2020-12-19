// Implementation of a simple ADT built around Go Arrays

package arrays

import (
	"fmt"
	"strings"
)

// Array definition
type Array struct {
	data []interface{}
}

// CreateArray will create a new array. First optional argument is the length
// and the second optional argument is the capacity. Panics if len is less than cap
func CreateArray(args ...uint) *Array {
	if args[0] > args[1] {
		panic("length cannot be greater than capacity")
	}
	arr := Array{}
	switch {
	case len(args) == 0:
		arr.data = make([]interface{}, 0)
	case len(args) == 1:
		arr.data = make([]interface{}, args[0])
	case len(args) == 2:
		arr.data = make([]interface{}, args[0], args[1])
	default:
		panic("Too many arguments provided")
	}

	return &arr
}

// Len returns the length of the underlying array
func (arr *Array) Len() int {
	return len(arr.data)
}

// Cap returns the capacity of the underlying array
func (arr *Array) Cap() int {
	return cap(arr.data)
}

// String implements the stringer interface
func (arr *Array) String() string {
	sb := make([]string, arr.Len())
	for i, val := range arr.data {
		sb[i] = fmt.Sprintf("%v", val)
	}
	return strings.Join(sb, ", ")
}
