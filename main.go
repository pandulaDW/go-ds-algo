package main

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/hashing"
)

func main() {
	arr := []*hashing.Data{{15, "foo"}, {32, "bar"}, {65, "bitch"},
		{3, "mat"}, {5, "food"}, {90, "rand"}, {42, "ryan"}}

	hashTable := hashing.NewHashTable(arr)
	fmt.Println(hashTable.SearchHashTable(15))

	hashTable.DeleteFromTable(82)
	fmt.Println(hashTable)
}
