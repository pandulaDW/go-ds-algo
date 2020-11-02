package main

import (
	"fmt"

	"algos.com/main/matrices"
)

func main() {
	d, err := matrices.CreateLowerTriangular([]int{2, 3, 5, 6, 8, 10, 4, 12, 11, 13}, -1)
	if err != nil {
		panic(err)
	}

	d.Set(2, 4, 49)

	fmt.Println(d)
}
