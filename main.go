package main

import (
	"fmt"

	"algos.com/main/linkedlist"
)

func main() {
	ll := linkedlist.CreateLinkedList()
	ll.Push(10)
	ll.Push(20)
	ll.Push(30)
	ll.Push(50)
	ll.Push(60)
	ll.Push(70)
	ll.Push(90)

	fmt.Println(ll.String())
}
