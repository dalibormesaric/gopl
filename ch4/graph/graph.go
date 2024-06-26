package main

import "fmt"

func main() {
	addEdge("a", "b")
	addEdge("b", "c")
	fmt.Println(hasEdge("a", "b"))
	fmt.Println(hasEdge("a", "c"))
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
