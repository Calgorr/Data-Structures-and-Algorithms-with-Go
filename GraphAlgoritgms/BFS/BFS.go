package main

import (
	"fmt"

	Queue "github.com/Calgorr/Data-Structures-and-Algorithms-with-Go/ElementaryDataStructures/queue/queue"
)

func main() {
	var Graph [][]int
	Graph = append(Graph, []int{1, 2})
	Graph = append(Graph, []int{1, 3, 4})
	Graph = append(Graph, []int{0, 3})
	Graph = append(Graph, []int{1, 2, 4})
	Graph = append(Graph, []int{1, 3})
	var visited [5]bool
	var rDist [5]int
	bfs(Graph, visited, rDist, 0)
}

func bfs(Graph [][]int, visited [5]bool, rDist [5]int, root int) {
	q := Queue.NewQueue[int]()
	q.Enqueue(root)
	visited[root] = true
	rDist[root] = 0
	for !q.IsEmpty() {
		node, _ := q.Dequeue()
		fmt.Printf("%v ", node)
		for _, v := range Graph[node] {
			if !visited[v] {
				visited[v] = true
				rDist[v] = rDist[node] + 1
				q.Enqueue(v)
			}
		}
	}
}
