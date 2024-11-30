package graph

import (
	"algorithm/helper"
)

func BFS(g *Graph, start int) {
	for _, u := range g.V {
		if u.ID != start {
			u.Visited = false
			u.In = int(^uint(0) >> 1)
			u.Predecessor = nil
		}
	}

	s := g.V[start]
	s.Visited = true
	s.In = 0
	s.Predecessor = nil
	Q := &helper.Queue[*Vertex]{}
	Q.Push(&s)

	for !Q.IsEmpty() {
		u := Q.Pop().(**Vertex)
		for _, v := range g.Adjency((*u).ID) {
			if !v.Visited {
				v.Visited = true
				v.In = (*u).In + 1
				v.Predecessor = (*u)
				Q.Push(&v)
			}
		}
	}
}
