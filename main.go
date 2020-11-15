package main

import (
	"fmt"
	"time"
)

func decorator(myFun func(name string) string) func(name string) string {
	inner := func(name string) string {
		start := time.Now()
		fmt.Println("Process started...")
		result := myFun(name)
		fmt.Printf("Process took %v\n", time.Since(start))
		return result
	}
	return inner
}

func main() {
	myFun1 := func(name string) string {
		time.Sleep(2 * time.Millisecond)
		return fmt.Sprintf("Running function named %s\n", name)
	}

	myFun2 := func(name string) string {
		time.Sleep(3 * time.Millisecond)
		return fmt.Sprintf("Running function named %s\n", name)
	}

	myFun1 = decorator(myFun1)
	myFun2 = decorator(myFun2)

	fmt.Println(myFun1("foo"))
	fmt.Println(myFun2("bar"))
}
