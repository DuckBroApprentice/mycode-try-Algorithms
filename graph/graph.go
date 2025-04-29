package graph

type Graph interface {
	AddNode(node int) error
	AddEdge(from, to int, weight int) error
	GetNeighbor(node int) ([]Edge, error)
	HasNode(node int) bool
	HasEdge(from, to int) bool
	IsDirected() bool
	IsWeight() bool
	NodeCount() int
	EdgeCount() int
	GetNodes() []int
	GetEdges(node int) ([]Edge, error)
}

type Edge struct {
	To     int
	Weight int
}
