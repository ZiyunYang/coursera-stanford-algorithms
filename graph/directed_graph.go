package graph

type DirGraph struct {
	graph
}

func NewDirected() *DirGraph {
	return &DirGraph{
		graph{
			edgesCount: 0,
			edges:      make(map[VertexId]map[VertexId]int),
			isDirected: true,
		},
	}
}

func (g *DirGraph) Reverse() *DirGraph {
	r := NewDirected()

	vertices := g.VerticesIter()
	for vertex := range vertices {
		r.AddVertex(vertex)
	}

	edges := g.EdgesIter()
	for edge := range edges {
		r.AddEdge(edge.To, edge.From, 1)
	}

	return r
}