package helpers

import (
	"math"
	"strconv"
)

// GetMaxElem returns the maximum element from an int slice
func GetMaxElem(arr []int) int {
	max := math.MinInt32
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

// LengthOfInt returns the width of a integer element
func LengthOfInt(n int) int {
	return len(strconv.FormatInt(int64(n), 10))
}

// SumOfNatural returns the sum of the given n, natural numbers
func SumOfNatural(n int) int {
	n++
	return n * (n - 1) / 2
}
