package graph

func DFS(g *Graph) {
	for _, u := range g.V {
		u.Predecessor = nil
		u.Visited = false
	}

	g.time = 0
	for _, u := range g.V {
		if !u.Visited {
			DfsVisit(g, u)
		}
	}
}

func DfsVisit(g *Graph, u *Vertex) {
	g.time += 1
	u.In = g.time
	u.Visited = true

	for _, v := range g.Adjency(u.ID) {
		if !v.Visited {
			v.Predecessor = u
			DfsVisit(g, v)
		}
	}

	g.time += 1
	u.Out = g.time
}
