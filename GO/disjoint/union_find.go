package disjoint

/* don't use this function.
 * use g.NewVertex */
func MakeSet(g *Graph, v *Vertex) {
	v.Parent = v
	v.Rank = 0
}

func FindSet(v *Vertex) *Vertex {
	if v != v.Parent {
		v.Parent = FindSet(v.Parent)
	}
	return v.Parent
}

func Union(g *Graph, f, t *Vertex) {
	rootF := FindSet(f)
	rootT := FindSet(t)
	if rootF != rootT {
		Link(rootF, rootT)
		g.NewEdge(f, t)
	}
}

func Link(x, y *Vertex) {
	if x.Rank > y.Rank {
		y.Parent = x
	} else {
		x.Parent = y
		if x.Rank == y.Rank {
			y.Rank = y.Rank + 1
		}
	}
}

/*
	Do not use this function.
*/
func ConnectedComponents(g *Graph) {
	for _, v := range g.V {
		MakeSet(g, v)
	}

	for from, edge := range g.E {
		f := g.V[from]
		for _, to := range edge {
			t := g.V[to]
			if FindSet(f) != FindSet(t) {
				Union(g, f, t)
			}
		}
	}
}
