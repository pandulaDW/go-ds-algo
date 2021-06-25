package main

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/hashing"
)

func main() {
	arr := []*hashing.Data{{8, "foo"}, {3, "bar"}, {6, "bitch"},
		{10, "mat"}, {15, "food"}, {9, "rand"}, {4, "ryan"}}

	hashTable := hashing.NewHashTable(arr)
	fmt.Println(hashTable.SearchHashTable(4))
}
