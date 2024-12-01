package mst

import (
	"algorithm/helper"
	"container/heap"
	"fmt"
)

func Prim(g *Graph, s *Vertex) []Edge {
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
	for i, v := range g.V {
		pq[i] = v
		v.index = i
	}
	heap.Init(&pq)

	result := []Edge{}
	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*Vertex)
		u.index = -1

		if u.Predecessor != nil {
			result = append(result, Edge{
				Name:   u.Predecessor.ID + "->" + u.ID,
				To:     u,
				Weight: g.getWeight(u.Predecessor, u),
			})
		}

		for _, e := range g.E[u] {
			v := e.To

			if v.index >= 0 && e.Weight < v.Distance {
				v.Predecessor = u
				v.Distance = e.Weight
				heap.Fix(&pq, v.index)
			}
		}

		// DEBUG
		fmt.Printf("current discover: %s \\ \n", u.ID)
		for i := range pq {
			fmt.Printf("%s's current cost: %d\n", pq[i].ID, pq[i].Distance)
		}
		fmt.Println("########################")
	}

	return result
}
