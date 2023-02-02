package main

import (
	"fmt"

	Stack "github.com/Calgorr/Data-Structures-and-Algorithms-with-Go/ElementaryDataStructures/stack/stack"
)

var stack *Stack.Stack[int]
var visited [5]int

func main() {
	graph := make([][]int, 5)
	graph[0] = []int{1, 3}
	graph[1] = []int{0, 3, 4}
	graph[2] = []int{0, 3}
	graph[3] = []int{1, 2, 4}
	graph[4] = []int{1, 3}
	stack = Stack.NewStack[int]()
	DFS(graph, 0)
}

func DFS(graph [][]int, root int) {
	stack.Push(root)
	visited[root] = 1
	for !stack.IsEmpty() {
		v1, _ := stack.Pop()
		fmt.Print(v1, " ")
		for _, v2 := range graph[v1] {
			if visited[v2] != 1 {
				visited[v2] = 1
				stack.Push(v2)
			}
		}
	}
	fmt.Println()
}
