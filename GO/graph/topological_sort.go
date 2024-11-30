package graph

import (
	"algorithm/sort"
)

type TopologicalSorter interface {
	Run(g *Graph) []int
}

type TopologicalSorterContext struct {
	Sorter TopologicalSorter
}

func (t *TopologicalSorterContext) SetSorter(sorter TopologicalSorter) {
	t.Sorter = sorter
}

func (t *TopologicalSorterContext) Run(g *Graph) []int {
	if t.Sorter == nil {
		panic("Sorter is not set")
	}
	return t.Sorter.Run(g)
}

type ArraySorter struct{}
type DFSSorter struct{}

func makeInDegree(g *Graph) map[int]int {
	inDegree := make(map[int]int)
	for _, v := range g.V {
		inDegree[v.ID] = 0
	}

	for _, e := range g.E {
		for _, v := range e {
			inDegree[v.ID]++
		}
	}

	return inDegree
}

func (s *ArraySorter) Run(g *Graph) []int {
	inDegree := makeInDegree(g)

	result := make([]int, 0)
	for len(result) != len(g.V) {
		for k, v := range inDegree {
			if v == 0 {
				inDegree[k] = -1

				adj := g.Adjency(k)
				for _, vertex := range adj {
					inDegree[vertex.ID]--
				}

				result = append(result, k)
			}
		}
	}

	return result
}

func (s *DFSSorter) Run(g *Graph) []int {
	DFS(g)

	beforeSort := make([]sort.Node[int], 0)
	for id, v := range g.V {
		beforeSort = append(beforeSort, sort.Node[int]{
			A: v.Out,
			B: id,
		})
	}

	sort.QuickSort(beforeSort, 0, len(beforeSort)-1)
	result := make([]int, 0)
	for i := 0; i < len(beforeSort); i++ {
		result = append(result, beforeSort[i].B)
	}

	return result
}
