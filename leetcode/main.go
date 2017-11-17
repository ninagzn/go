package main

import "fmt"

func main() {
	minStack := MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin()

	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())

}
