package main

import "fmt"

func main() {
	stack := NewStack[int]()
	stack.push(1)
	stack.push(2)
	fmt.Println(stack.Top())
	stack.push(3)
	stack.push(4)
	fmt.Println(stack.Top())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
