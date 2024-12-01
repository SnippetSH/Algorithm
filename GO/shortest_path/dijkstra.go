package shortestpath

import (
	"algorithm/helper"
	"container/heap"
)

// 방향 그래프
type Vertex struct {
	ID       string
	Distance int
	index    int
}

func (v *Vertex) Cmp() int {
	return v.Distance
}

func (v *Vertex) Index() *int {
	return &v.index
}

func (v *Vertex) SetIndex(i int) {
	v.index = i
}

type Edge struct {
	Weight int
	To     *Vertex
	From   *Vertex
}

type Graph struct {
	V map[string]*Vertex
	E map[*Vertex][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		V: make(map[string]*Vertex),
		E: make(map[*Vertex][]*Edge),
	}
}

func (g *Graph) NewVertex(id string) *Vertex {
	if g.V[id] != nil {
		return nil
	}

	v := &Vertex{
		ID:       id,
		Distance: int(^uint32(0) >> 1),
	}
	g.V[id] = v
	return v
}

func (g *Graph) NewEdge(f, t *Vertex, w int) {
	for _, e := range g.E[f] {
		if e.To == t && e.From == f {
			return
		}
	}

	g.E[f] = append(g.E[f], &Edge{
		To:     t,
		From:   f,
		Weight: w,
	})
}

func Dijkstra(g *Graph, s *Vertex) {
	flag := false
	for _, v := range g.V {
		if v == s {
			flag = true
			break
		}
	}
	if !flag {
		panic("start vertex is not in graph")
	}

	s.Distance = 0
	pq := make(helper.PriorityQueue[*Vertex], len(g.V))
	idx := 0
	for _, v := range g.V {
		pq[idx] = v
		v.index = idx
		idx++
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Vertex)

		for _, e := range g.E[u] {
			if g.V[e.To.ID].Distance > u.Distance+e.Weight {
				g.V[e.To.ID].Distance = u.Distance + e.Weight
				heap.Fix(&pq, g.V[e.To.ID].index)
			}
		}
	}
}
