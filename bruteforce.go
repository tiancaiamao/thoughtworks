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

func BruteForceDistance(g *Graph, dist int, start byte, end byte) int {
	v1, ok := g.Name2Vertex[start]
	if !ok {
		return 0
	}
	v2, ok := g.Name2Vertex[end]
	if !ok {
		return 0
	}

	return bruteForceDistanceRecur(g, dist, v1, v2, true)
}

func bruteForceDistanceRecur(g *Graph, dist int, v1 int, v2 int, self bool) int {
	if dist <= 0 {
		return 0
	}

	ret := 0
	if !self {
		if v1 == v2 {
			ret++
		}
	}

	for e := g.Edges[v1]; e != nil; e = e.Next {
		ret += bruteForceDistanceRecur(g, dist-int(e.Weight), e.ToVertex, v2, false)
	}
	return ret
}
