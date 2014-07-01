package main

import (
	// "fmt"
)

type Edge struct {
	ToVertex int   // adjacency info
	Next     *Edge // next edge in list
	Weight   uint
}

type Graph struct {
	Edges       []*Edge      // adjacency info
	Vertexs       []byte       // for DEBUG ONLY, not needed actually
	Name2Vertex map[byte]int // map Node name to Vertex
}

func NewGraph() *Graph {
	ret := new(Graph)
	ret.Edges = make([]*Edge, 0, 10)
	ret.Vertexs = make([]byte, 0, 10)
	ret.Name2Vertex = make(map[byte]int)
	return ret
}

func (g *Graph) addVertex(name byte) {
	if _, ok := g.Name2Vertex[name]; !ok {
		g.Name2Vertex[name] = len(g.Edges)
		g.Edges = append(g.Edges, nil)
		g.Vertexs = append(g.Vertexs, name)
	}
}

// internal api, use index for vertex
func (g *Graph) addEdge(from int, to int, weight uint) {
	head := g.Edges[from]
	e := &Edge{
		ToVertex: to,
		Weight:   weight,
	}
	if head == nil {
		g.Edges[from] = e
	} else {
		e.Next = head.Next
		head.Next = e
	}
}

// exported api, use name for vertex
func (g *Graph) AddEdge(from byte, to byte, weight uint) {
	vert1, ok := g.Name2Vertex[from]
	if !ok {
		g.addVertex(from)
		vert1 = g.Name2Vertex[from]
	}

	vert2, ok := g.Name2Vertex[to]
	if !ok {
		g.addVertex(to)
		vert2 = g.Name2Vertex[to]
	}

	g.addEdge(vert1, vert2, weight)
}



func main() {
}
