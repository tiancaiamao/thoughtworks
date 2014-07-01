package main

import (
	"fmt"
	"testing"
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

func makeGraph() *Graph {
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

	return g
}

func TestGraph(t *testing.T) {
	g := makeGraph()
	DFS(g)
}

func TestDijkstra(t *testing.T) {
	g := makeGraph()

	testDijkstra('A', 'B', 5, g, t)
	testDijkstra('B', 'A', UINT_MAX, g, t)
	testDijkstra('E', 'B', 3, g, t)
	testDijkstra('E', 'C', 7, g, t)
	testDijkstra('A', 'A', UINT_MAX, g, t)
	testDijkstra('B', 'B', 9, g, t)
}

func testDijkstra(start byte, to byte, value uint, g *Graph, t *testing.T) {
	if g.Dijkstra(start, to) != value {
		t.Errorf("dijkstra...(%c, %c)\n", start, to)
	}
}
