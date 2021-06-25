package sorting

func predicate(i, j int, asc bool) bool {
	if asc {
		return i > j
	}
	return i < j
}

// BubbleSort will sort the passed array using bubble sort method
func BubbleSort(arr []int, asc bool) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	comparisonsNeeded := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		isSorted := true
		for j := 0; j < comparisonsNeeded; j++ {
			if predicate(sorted[j], sorted[j+1], asc) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
		comparisonsNeeded--
	}

	return sorted
}
