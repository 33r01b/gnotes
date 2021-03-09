package search

// UnweightedGraph is
type UnweightedGraph struct {
	edges map[Node][]*Node
}

// NewUnweightedGraph creates a new UnweightedGraph
func NewUnweightedGraph() *UnweightedGraph {
	return &UnweightedGraph{
		edges: make(map[Node][]*Node),
	}
}

// AddEdges adds an edge to graph
func (g *UnweightedGraph) AddEdges(n *Node, nodes []*Node) {
	g.edges[*n] = nodes
}

// Edges returns Node edges
func (g *UnweightedGraph) Edges(n *Node) (edges []*Node, ok bool) {
	edges, ok = g.edges[*n]
	return
}

// Node is
type Node struct {
	name    string
	profile string
}
