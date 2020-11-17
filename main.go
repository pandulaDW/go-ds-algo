package main

import (
	"fmt"

	"algos.com/main/customio"
)

func main() {
	store := customio.CreateStore(10)

	s := "Hello, World!!"
	store.Write([]byte(s))

	b1 := make([]byte, 5)
	store.Read(b1)
	fmt.Println(string(b1))

	b2 := make([]byte, 10)
	store.Flush(b2)
	fmt.Println(string(b2))
}
