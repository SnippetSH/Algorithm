package shortestpath

func BellmanFord(g *Graph, s *Vertex) bool {
	s.Distance = 0

	for i := 0; i < len(g.V)-1; i++ {
		for _, E := range g.E {
			for _, e := range E {
				if e.To.Distance > e.From.Distance+e.Weight {
					e.To.Distance = e.From.Distance + e.Weight
				}
			}
		}
	}

	for _, E := range g.E {
		for _, e := range E {
			if e.To.Distance > e.From.Distance+e.Weight {
				return false
			}
		}
	}

	return true
}

func MoreEfficientBellmanFord(g *Graph, s *Vertex) bool {
	s.Distance = 0

	for i := 0; i < len(g.V)-1; i++ {
		updated := false
		for _, E := range g.E {
			for _, e := range E {
				if e.To.Distance > e.From.Distance+e.Weight {
					e.To.Distance = e.From.Distance + e.Weight
					updated = true
				}
			}
		}

		if !updated {
			break
		}
	}

	for _, E := range g.E {
		for _, e := range E {
			if e.To.Distance > e.From.Distance+e.Weight {
				return false
			}
		}
	}

	return true
}
