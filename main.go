package main

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/hashing"
	"unicode/utf8"
)

func main() {
	arr := []*hashing.DataWithNumID{{15, "foo"}, {32, "bar"}, {65, "bitch"},
		{3, "mat"}, {5, "food"}, {90, "rand"}, {42, "ryan"}}

	hashTable := hashing.NewHashTableForNumID(arr)
	fmt.Println(hashTable.SearchHashTable(15))

	hashTable.DeleteFromTable(82)
	fmt.Println(hashTable)

	decodeRune, _ := utf8.DecodeRune([]byte("A"))
	fmt.Println(decodeRune)
}
