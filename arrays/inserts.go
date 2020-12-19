package arrays

// Push will push elements to the array at the end returns the struct
func (arr *Array) Push(el interface{}) *Array {
	arr.data = append(arr.data, el)
	return arr
}
