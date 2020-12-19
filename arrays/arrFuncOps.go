package arrays

// Predicate function type def
type Predicate func(element interface{}, index int) bool

// Mapper function type def
type Mapper func(element interface{}, index int) interface{}

// Filter will filter only the suitable elements in the underlying array
func (arr *Array) Filter(test Predicate) *Array {
	filtered := make([]interface{}, 0)

	for i, val := range arr.data {
		if test(val, i) {
			filtered = append(filtered, val)
		}
	}

	arr.data = filtered
	return arr
}

// Map will modify each element in the array according the mapper function
func (arr *Array) Map(mapper Mapper) *Array {
	for i, val := range arr.data {
		arr.data[i] = mapper(val, i)
	}
	return arr
}
