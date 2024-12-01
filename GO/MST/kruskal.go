package mst

import "algorithm/sort"

type KVertex struct {
	ID     string
	Cost   int
	Parent *KVertex
	Rank   int
}

type KEdge struct {
	Weight int
	To     *KVertex
	From   *KVertex
}

type KGraph struct {
	V []*KVertex
	E []*KEdge
}

func KNewGraph() *KGraph {
	return &KGraph{
		V: make([]*KVertex, 0),
		E: make([]*KEdge, 0),
	}
}

func (g *KGraph) NewVertex(id string) *KVertex {
	v := &KVertex{
		ID:   id,
		Cost: 0,
		Rank: 0,
	}
	v.Parent = v
	g.V = append(g.V, v)

	return v
}

func (g *KGraph) NewEdge(f, t *KVertex, w int) {
	for _, e := range g.E {
		if e.To == t && e.From == f {
			return
		}
	}

	g.E = append(g.E, &KEdge{
		From:   f,
		To:     t,
		Weight: w,
	})
}

func (g *KGraph) getWeight(f, t *KVertex) int {
	for _, v := range g.E {
		if v.To == t && v.From == f {
			return v.Weight
		}
	}

	panic("That vertex is not connected")
}

func FindSet(v *KVertex) *KVertex {
	if v != v.Parent {
		v.Parent = FindSet(v.Parent)
	}

	return v.Parent
}

func Union(f, t *KVertex) {
	rootF := FindSet(f)
	rootT := FindSet(t)
	Link(rootF, rootT)
}

func Link(x, y *KVertex) {
	if x.Rank > y.Rank {
		y.Parent = x
	} else {
		x.Parent = y
		if x.Rank == y.Rank {
			y.Rank = y.Rank + 1
		}
	}
}

func Kruskal(g *KGraph) []Edge {
	arr := make([](sort.Node[*KEdge]), 0)
	for _, e := range g.E {
		arr = append(arr, sort.Node[*KEdge]{
			A: e.Weight,
			B: e,
		})
	}

	sort.QuickSort(arr, 0, len(arr)-1)
	for i := 0; i < len(arr); i++ {
		g.E[i] = arr[i].B
	}

	result := make([]Edge, 0)
	for _, e := range g.E {
		u := e.From
		v := e.To
		if FindSet(u) != FindSet(v) {
			Union(u, v)

			result = append(result, Edge{
				Name:   u.ID + "->" + v.ID,
				Weight: e.Weight,
			})
		}
	}

	return result
}
