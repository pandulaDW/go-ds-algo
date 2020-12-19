// Implementation of a simple ADT built around Go Arrays

package arrays

// Array definition
type Array struct {
	data []interface{}
}

// Len returns the length of the underlying array
func (arr *Array) Len() int {
	return len(arr.data)
}

// Cap returns the capacity of the underlying array
func (arr *Array) Cap() int {
	return cap(arr.data)
}
