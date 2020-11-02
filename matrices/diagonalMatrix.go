package matrices

import (
	"errors"
	"fmt"

	"algos.com/main/helpers"
)

// Diagonal matrices -> since only diagonal have non zero elements
// we can just use one dimension array to store the diagonal elements
type Diagonal struct {
	data []int
	ndim int
}

// CreateDiagonal creates a diagonal matrix
func CreateDiagonal(n int) Diagonal {
	d := Diagonal{}
	d.data = make([]int, n)
	d.ndim = n
	return d
}

// Set val into index position given by i and j
func (d *Diagonal) Set(i, j, val int) error {

	if i >= d.ndim || j >= d.ndim {
		return errors.New(`row and column indices cannot exceed the matrix dimensions`)
	}

	if i != j {
		return errors.New(`Diagonal matrices can only create elements within it's Diagonal line`)
	}

	d.data[i] = val
	return nil
}

// Get value from index position given by i and j
func (d *Diagonal) Get(i, j int) (int, error) {
	if i >= d.ndim || j >= d.ndim {
		return 0, errors.New(`row and column indices cannot exceed the matrix dimensions`)
	}

	if i != j {
		return 0, nil
	}

	return d.data[i], nil
}

// Display matrix like a dataframe
func (d *Diagonal) Display() {
	maxElem := helpers.GetMaxElem(d.data)
	maxWidth := helpers.LengthOfInt(maxElem)

	for i := 0; i < d.ndim; i++ {
		for j := 0; j < d.ndim; j++ {
			if i == j {
				fmt.Print(d.data[i])
				padding := maxWidth - helpers.LengthOfInt(d.data[i]) + 1
				fmt.Printf("%*s", padding, "|")
			} else {
				fmt.Print(0)
				fmt.Printf("%*s", maxWidth, "|")
			}
		}
		fmt.Println()
	}
}
