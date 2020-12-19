package main

import (
	"fmt"

	"algos.com/main/arrays"
)

func main() {
	arr := arrays.CreateArray(0, 10)
	arr.Push(12).Push(23).Push(32).Push(41)

	arr.Filter(func(el interface{}, idx int) bool {
		return el.(int)%2 == 0
	}).Map(func(el interface{}, idx int) interface{} {
		return el.(int) * el.(int)
	})

	fmt.Println(arr.String())
}
