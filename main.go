package main

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/hashing"
	"github.com/pandulaDW/go-ds-algo/linkedList"
)

func main() {
	arr := []*hashing.Data{{8, "foo"}, {3, "bar"}, {6, "bitch"},
		{10, "mat"}, {15, "food"}, {9, "rand"}, {4, "ryan"}}

	hashTable := hashing.NewHashTable(arr)
	fmt.Println(hashTable.SearchHashTable(4))

	list := linkedList.NewLinkedList()
	list.InsertAtStart(100)
	list.InsertAtEnd(12)
	list.InsertAtEnd(15)
	list.InsertAtEnd(32)

	fmt.Println(list)
}
