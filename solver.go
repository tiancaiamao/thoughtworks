package main

import "fmt"

type Problem interface {
	Solve(g *Graph)
}

type Quest4 struct {
	id   int
	from byte
	to   byte
	dist int
}

func (prob *Quest4) Solve(g *Graph) {
	n := BruteForceDistance(g, prob.dist, prob.from, prob.to)
	fmt.Printf("Output #%d: %d\n", prob.id, n)
}

type Quest2 struct {
	id      int
	start   byte
	end     byte
	exactly bool // exactly or maximum
	stops   int
}

type Quest1 struct {
	id    int
	input []byte
}

func (prob *Quest1) Solve(g *Graph) {
	input := prob.input
	start := input[0]
	current, ok := g.Name2Vertex[start]
	if !ok {
		fmt.Printf("I have no idea what you are talking about")
		return
	}

	result := uint(0)
	for i := 1; i < len(input); i++ {
		name := input[i]
		next, ok := g.Name2Vertex[name]
		if !ok {
			fmt.Printf("I have no idea what you are talking about")
			return
		}

		find := false
		for e := g.Edges[current]; e != nil; e = e.Next {
			if e.ToVertex == next {
				find = true
				result += e.Weight
				current = next
				break
			}
		}

		if !find {
			fmt.Printf("Output #%d: NO SUCH ROUTE\n", prob.id)
			return
		}
	}
	fmt.Printf("Output #%d: %d\n", prob.id, result)
}

func (prob *Quest2) Solve(g *Graph) {
	within, exactly := BruteForceNStops(g, prob.stops, prob.start, prob.end)
	if prob.exactly {
		fmt.Printf("Output #%d: %d\n", prob.id, exactly)
	} else {
		fmt.Printf("Output #%d: %d\n", prob.id, within)
	}
}

type Quest3 struct {
	id   int
	from byte
	to   byte
}

func (prob *Quest3) Solve(g *Graph) {
	dist := g.Dijkstra(prob.from, prob.to)
	if dist == UINT_MAX {
		fmt.Printf("Output #%d: NO SUCH ROUTE\n", prob.id)
	} else {
		fmt.Printf("Output #%d: %d\n", prob.id, dist)
	}
}