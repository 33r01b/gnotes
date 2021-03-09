package search

// Graph is
type Graph struct {
	edges map[Node][]*Node
}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{
		edges: make(map[Node][]*Node),
	}
}

// AddEdges adds an edge to graph
func (g *Graph) AddEdges(n *Node, nodes []*Node) {
	g.edges[*n] = nodes
}

// Edges returns Node edges
func (g *Graph) Edges(n *Node) (edges []*Node, ok bool) {
	edges, ok = g.edges[*n]
	return
}

// Node is
type Node struct {
	name    string
	profile string
}
