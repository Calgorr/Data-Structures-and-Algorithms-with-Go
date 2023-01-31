package main

import "fmt"

var time int

func main() {
	graph := make([][]int, 5)
	graph[0] = []int{1, 3}
	graph[1] = []int{1, 3, 4}
	graph[2] = []int{0, 3}
	graph[3] = []int{1, 2, 4}
	graph[4] = []int{1, 3}
	var visited, parent, t1, t2 [5]int
	dfs(graph, visited, parent, t1, t2)

}

func dfs(graph [][]int, visited, parent, t1, t2 [5]int) {
	for k := range graph {
		if visited[k] == 0 {
			dfsVisit(graph, visited, parent, t1, t2, k)
		}
	}
}

func dfsVisit(graph [][]int, visited, parent, t1, t2 [5]int, root int) {
	time++
	t1[root] = time
	visited[root] = 1
	for k2, v2 := range graph[root] {
		if visited[v2] == 0 {
			fmt.Println(graph[root][k2])
			visited[v2] = 1
			parent[v2] = root
			dfsVisit(graph, visited, parent, t1, t2, root)

		}
	}
	visited[root] = 2
	time++
	t2[root] = time
}
