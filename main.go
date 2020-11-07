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
	ll.Push(40)
	ll.Push(60)
	ll.Push(60)
	ll.Push(80)

	ll2 := linkedlist.CreateLinkedList()
	ll2.Push(5)
	ll2.Push(32)
	ll2.Push(64)
	ll2.Push(82)

	ll.Merge(&ll2)
	fmt.Println(ll.String())
}
