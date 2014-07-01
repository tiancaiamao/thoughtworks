package main

func BruteForceNStops(g *Graph, N int, start byte, end byte) (int, int) {
	within := 0
	exactly := 0
	v1, ok := g.Name2Vertex[start]
	if !ok {
		return 0, 0
	}
	v2, ok := g.Name2Vertex[end]
	if !ok {
		return 0, 0
	}

	bruteForceNStopsRecur(g, 0, N, v1, v2, &within, &exactly)
	return within, exactly
}

func bruteForceNStopsRecur(g *Graph, n int, N int, v1 int, v2 int, within *int, exactly *int) {
	if n > N {
		return
	}

	if n > 0 && v1 == v2 {
		(*within)++
		if n == N {
			(*exactly)++
		}
	}

	n++
	for e := g.Edges[v1]; e != nil; e = e.Next {
		bruteForceNStopsRecur(g, n, N, e.ToVertex, v2, within, exactly)
	}
}
