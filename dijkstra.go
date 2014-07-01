package main

const UINT_MAX uint = 9999999

func (g *Graph) Dijkstra(s, t byte) uint {
	v1, ok := g.Name2Vertex[s]
	if !ok {
		return UINT_MAX
	}

	v2, ok := g.Name2Vertex[t]
	if !ok {
		return UINT_MAX
	}

	known := make(map[int]struct{})

	// init the shortest path dist
	dist := make([]uint, len(g.Edges))
	for i := 0; i < len(dist); i++ {
		dist[i] = UINT_MAX
	}

	// start from the start vertex
	for e := g.Edges[v1]; e != nil; e = e.Next {
		dist[e.ToVertex] = e.Weight
	}

	last := -1

	for last != v2 {
		// select next vertex from the unknown vertex minimizing dist
		vNext, _ := min(dist, known)
		if vNext == -1 {
			// error, can't find a path from s to t
			return UINT_MAX
		}

		for e := g.Edges[vNext]; e != nil; e = e.Next {
			if dist[vNext]+e.Weight < dist[e.ToVertex] {
				dist[e.ToVertex] = dist[vNext] + e.Weight
			}
		}

		last = vNext
		known[vNext] = struct{}{}
	}
	return dist[v2]
}

func min(dist []uint, known map[int]struct{}) (int, uint) {
	minVal := UINT_MAX
	minIdx := -1
	for i := 0; i < len(dist); i++ {
		if _, ok := known[i]; !ok {
			if dist[i] < minVal {
				minVal = dist[i]
				minIdx = i
			}
		}
	}
	return minIdx, minVal
}
