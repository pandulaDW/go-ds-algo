package main

import (
	"fmt"
	"github.com/pandulaDW/go-ds-algo/data"
	"github.com/pandulaDW/go-ds-algo/hashing"
)

func main() {
	arr1 := []*hashing.DataWithNumID{{15, "foo"}, {32, "bar"}, {65, "bitch"},
		{3, "mat"}, {5, "food"}, {90, "rand"}, {42, "ryan"}}

	table1 := hashing.NewHashTableForNumID(arr1)
	fmt.Println(table1.SearchHashTable(65))

	accounts := data.ReadAccountData()

	hashItems := data.ConvertToHashTableType(accounts)
	table2 := hashing.NewHashTableForStringID(hashItems[0:30])
	fmt.Println(table2.SearchHashTable("5ca4bbc7a2dd94ee58162486"))
}
