package disjoint

import (
	"fmt"
	"strings"
)

type id int
type Vertex struct {
	ID     id
	Parent *Vertex
	Rank   int
}

type Graph struct {
	V      map[id]*Vertex
	E      map[id][]id
	nextID id
}

func NewGraph() *Graph {
	return &Graph{
		V: make(map[id]*Vertex),
		E: make(map[id][]id),
	}
}

func (g *Graph) NewVertex() *Vertex {
	v := &Vertex{
		ID:   g.nextID,
		Rank: 0,
	}
	v.Parent = v
	g.V[g.nextID] = v
	g.nextID++
	return v
}

func (g *Graph) NewEdge(f, t *Vertex) {
	for _, v := range g.E[f.ID] {
		if v == t.ID {
			return
		}
	}
	g.E[f.ID] = append(g.E[f.ID], t.ID)
	// g.E[t.ID] = append(g.E[t.ID], f.ID)
}

func (g *Graph) String() string {
	var sb strings.Builder
	for id, v := range g.V {
		sb.WriteString(fmt.Sprintf("Vertex %d -> ", id))
		for _, edge := range g.E[v.ID] {
			sb.WriteString(fmt.Sprintf("%d ", edge))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
