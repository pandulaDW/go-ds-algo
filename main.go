package main

import (
	"fmt"

	"algos.com/main/stacks"
)

func main() {
	stack := stacks.CreateStackUsingArray(10)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	stack.Pop()

	fmt.Println(stack.String())
}
