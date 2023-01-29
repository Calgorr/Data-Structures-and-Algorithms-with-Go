package main

import (
	"fmt"
	Queue "main/queue"
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
	q := Queue.NewQueue[int]()
	q.Enqueue(0)
	visited[0] = true
	rDist[0] = 0
	bfs(Graph, visited, rDist, 0)
}

func bfs(Graph [][]int, visited [5]bool, rDist [5]int, q Queue.Queue[int]) {
	if q.IsEmpty() {
		return
	}
	node, _ := q.Dequeue()
	fmt.Printf("%v ", node)
	for _, v := range Graph[node] {
		if !visited[v] {
			visited[v] = true
			rDist[v] = rDist[node] + 1
			q.Enqueue(v)
		}
	}
	bfs(Graph, visited, rDist, q)
}
