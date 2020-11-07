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

	loopList := linkedlist.DummyLoop()
	fmt.Println(loopList.IsLoop())

	fmt.Println(ll.String())
}
