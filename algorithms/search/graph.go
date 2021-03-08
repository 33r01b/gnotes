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

// AddEdge adds an edge to graph
func (g *Graph) AddEdge(n *Node, nodes []*Node) {
	g.edges[*n] = nodes
}

// Node is
type Node struct {
	name    string
	profile string
}
