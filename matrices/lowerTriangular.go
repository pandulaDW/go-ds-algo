package matrices

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"algos.com/main/helpers"
)

// LowerTriangular matrices has non-zero elements below and inclusive
// of the diagonal.
//
// Meaning that if i < j, arr[i, j] should be zero and
// if i >= j, arr[i, j] should be non-zero
type LowerTriangular struct {
	data []int
	ndim int
}

// CreateLowerTriangular creates a CreateLowerTriangular struct with the given array
//
// if n is set as -1, the function will try to infer the dimension, if it's not
// -1 and its a positive dimension, the function will check its validity. If the provided
// dimension or array size are are not valid, then an error will be thrown
func CreateLowerTriangular(initArr []int, n int) (LowerTriangular, error) {

	d := LowerTriangular{}

	if n == -1 {
		for i := 1; i < len(initArr); i++ {
			if len(initArr) == helpers.SumOfNatural(i) {
				d.ndim = i
				d.data = initArr
				return d, nil
			}

			if len(initArr) < helpers.SumOfNatural(i) {
				return LowerTriangular{}, errors.New("Array length is not a valid")
			}
		}
	}

	if n <= 0 && n != -1 {
		return LowerTriangular{}, errors.New("Array dimensions cannot be negative or zero")
	}

	if n > 0 {
		if len(initArr) != helpers.SumOfNatural(n) {
			return LowerTriangular{}, errors.New("The given dimension is incorrect")
		}
		d.ndim = n
	}

	arrSize := helpers.SumOfNatural(n)
	d.data = make([]int, arrSize)
	return d, nil
}

// Set the given value at the index positions using a row major formula
func (d *LowerTriangular) Set(i, j, val int) error {
	if i >= d.ndim || j >= d.ndim {
		return errors.New(`row and column indices cannot exceed the matrix dimensions`)
	}

	if i < j {
		return errors.New(`LowerTriangular matrices cannot have non zero elements 
		above the diagonal line`)
	}

	idxPosition := helpers.SumOfNatural(i-1) + j - 1
	d.data[idxPosition] = val
	return nil
}

// Get value from index position given by i and j
func (d *LowerTriangular) Get(i, j int) (int, error) {
	if i >= d.ndim || j >= d.ndim {
		return 0, errors.New(`row and column indices cannot exceed the matrix dimensions`)
	}

	if i < j {
		return 0, nil
	}

	idxPosition := helpers.SumOfNatural(i-1) + j - 1
	return d.data[idxPosition], nil
}

// String returns a string representation of the matrix
func (d LowerTriangular) String() string {
	var sb strings.Builder

	genereteRow := func(rowIndex int) string {
		var tempSb strings.Builder
		startIndex := helpers.SumOfNatural(rowIndex - 1)
		endIndex := helpers.SumOfNatural(rowIndex)
		nonZeroElements := d.data[startIndex:endIndex]

		for i := 0; i < d.ndim; i++ {
			if i < len(nonZeroElements) {
				elem := strconv.FormatInt(int64(nonZeroElements[i]), 10)
				tempSb.WriteString(elem + ", ")
			} else {
				tempSb.WriteString("0, ")
			}
		}
		return tempSb.String()
	}

	for i := 1; i < d.ndim+1; i++ {
		row := genereteRow(i) + "\n"
		sb.WriteString(row)
	}

	footer := fmt.Sprintf("\nsize: %dx%d\n", d.ndim, d.ndim)
	sb.WriteString(footer)
	return sb.String()
}
