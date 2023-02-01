package main

import "fmt"

var time int
var visited, parent, t1, t2 [5]int

func main() {
	graph := make([][]int, 5)
	graph[0] = []int{1, 3}
	graph[1] = []int{0, 3, 4}
	graph[2] = []int{0, 3}
	graph[3] = []int{1, 2, 4}
	graph[4] = []int{1, 3}
	fmt.Printf("0 ")
	dfs(graph)

}

func dfs(graph [][]int) {
	for k := range graph {
		if visited[k] == 0 {
			dfsVisit(graph, k)
		}
	}
}

func dfsVisit(graph [][]int, root int) {
	time++
	t1[root] = time
	visited[root] = 1
	for _, v2 := range graph[root] {
		if visited[v2] == 0 {
			fmt.Printf("%v ", v2)
			parent[v2] = root
			dfsVisit(graph, v2)

		}
	}
	visited[root] = 2
	time++
	t2[root] = time
}
