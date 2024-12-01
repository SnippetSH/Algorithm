package mst

// 무방향 그래프
type Vertex struct {
	ID          string
	Distance    int
	Predecessor *Vertex
	Visited     bool
	index       int
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
	Name   string
	To     *Vertex
	Weight int
}

type Graph struct {
	V []*Vertex
	E map[*Vertex][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		V: make([]*Vertex, 0),
		E: make(map[*Vertex][]*Edge),
	}
}

func (g *Graph) NewVertex(id string) *Vertex {
	v := &Vertex{
		ID:          id,
		Distance:    int(^uint32(0) >> 1),
		Predecessor: nil,
		Visited:     false,
	}
	g.V = append(g.V, v)
	return v
}

func (g *Graph) NewEdge(f, t *Vertex, w int) {
	for _, e := range g.E[f] {
		if e.To == t {
			return
		}
	}

	g.E[f] = append(g.E[f], &Edge{
		Name:   f.ID + t.ID,
		To:     t,
		Weight: w,
	})
	g.E[t] = append(g.E[t], &Edge{
		Name:   t.ID + f.ID,
		To:     f,
		Weight: w,
	})
}

func (g *Graph) getWeight(f, t *Vertex) int {
	for _, v := range g.E[f] {
		if v.To == t {
			return v.Weight
		}
	}

	panic("That vertex is not connected")
}
