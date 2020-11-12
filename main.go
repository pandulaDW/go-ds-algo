package main

import (
	"fmt"

	"algos.com/main/stacks"
)

func main() {
	stack := stacks.CreateStackUsingList(10)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	// stack.Pop()

	fmt.Println("First element : ", stack.Peek(0))
	fmt.Println("Fourth element : ", stack.Peek(3))
	fmt.Println("Last element : ", stack.Peek(4))

	fmt.Println(stack.String())
}
