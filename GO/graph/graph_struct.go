package graph

type Vertex struct {
	ID          int
	In          int // d
	Out         int // f
	Visited     bool
	Predecessor *Vertex
}

type Graph struct {
	V    map[int]*Vertex
	E    map[int][]*Vertex
	time int
}

func NewGraph() *Graph {
	return &Graph{
		V:    make(map[int]*Vertex),
		E:    make(map[int][]*Vertex),
		time: 0,
	}
}

func (g *Graph) Adjency(ID int) []*Vertex {
	return g.E[ID]
}

func (g *Graph) AddVertex(id int) {
	if _, exists := g.V[id]; !exists {
		g.V[id] = &Vertex{ID: id}
	}
}

func (g *Graph) AddEdge(from, to int) {
	if _, exists := g.V[from]; !exists {
		g.AddVertex(from)
	}
	if _, exists := g.V[to]; !exists {
		g.AddVertex(to)
	}

	g.E[from] = append(g.E[from], g.V[to])
}
