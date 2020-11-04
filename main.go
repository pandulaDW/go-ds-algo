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
	ll.Push(50)
	ll.Push(110)
	ll.Push(90)

	fmt.Println(ll.String())

	fmt.Println(ll.FindByDataRepeated(90))
	ll.Push(120)

	fmt.Println(ll.String())
}
