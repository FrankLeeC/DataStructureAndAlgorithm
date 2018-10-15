package graph

// Vertex vertex of graph
type Vertex interface {
	Value() *interface{}
	UniqueID() string
}

// Edge edge of graph
type Edge interface {
	Weight() *interface{}
	Start() Vertex
	End() Vertex
}

// Graph graph
type Graph interface {
	AddEdge(e Edge)
	Adj(v Vertex) []Vertex
	ShortestPath(s, t Vertex) ([]Vertex, []interface{})
}
