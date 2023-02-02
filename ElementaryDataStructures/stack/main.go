package main

import (
	"fmt"

	Stack "github.com/Calgorr/Data-Structures-and-Algorithms-with-Go/ElementaryDataStructures/stack/stack"
)

func main() {
	stack := Stack.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	stack.Push(3)
	stack.Push(4)

	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
