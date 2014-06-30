package main

import (
	"testing"
	"fmt"
)

func DFS(g *Graph) {
	visited := make([]bool, len(g.Edges))

	dfs(g, 0, visited)
}

func dfs(g *Graph, v int, visited []bool) {
	if visited[v] {
		return
	}

	fmt.Printf("%c", g.Nodes[v])
	visited[v] = true

	for e := g.Edges[v]; e != nil; e = e.Next {
		dfs(g, e.ToVertex, visited)
	}
}

func TestGraph(t *testing.T) {
	g := NewGraph()
	g.AddEdge('A', 'B', 5)
	g.AddEdge('B', 'C', 4)
	g.AddEdge('C', 'D', 8)
	g.AddEdge('D', 'C', 8)
	g.AddEdge('D', 'E', 6)
	g.AddEdge('A', 'D', 5)
	g.AddEdge('C', 'E', 2)
	g.AddEdge('E', 'B', 3)
	g.AddEdge('A', 'E', 7)

	DFS(g)
}
