package main

// O(2n) time
func findDuplicateElements(arr []int) []int {
	valueCounts := map[int]int{}

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if _, ok := valueCounts[val]; !ok {
			valueCounts[val] = 1
		} else {
			valueCounts[val]++
		}
	}

	duplicates := make([]int, 0, len(valueCounts))
	for key, val := range valueCounts {
		if val > 1 {
			duplicates = append(duplicates, key)
		}
	}
	return duplicates
}
