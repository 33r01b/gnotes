package search

// Node is
type Node struct {
	Name    string
	Profile string
}

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

// AddEdges adds an edges to graph
func (g *UnweightedGraph) AddEdges(n *Node, nodes []*Node) {
	g.edges[*n] = nodes
}

// Edges returns Node edges
func (g *UnweightedGraph) Edges(n *Node) (edges []*Node, ok bool) {
	edges, ok = g.edges[*n]
	return
}

// WeightedGraph is
type WeightedGraph struct {
	// from -> to -> weight
	edges map[Node]map[Node]int
}

// NewWeightedGraph creates a new WeightedGraph
func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		edges: make(map[Node]map[Node]int),
	}
}

// AddEdge adds an edge to graph
// set to as nil, if from is final node
func (g *WeightedGraph) AddEdge(from *Node, to *Node, weight int) {
	if _, ok := g.edges[*from]; !ok {
		g.edges[*from] = make(map[Node]int)
	}

	if to != nil {
		g.edges[*from][*to] = weight
	}
}

// Edges returns Node edges
func (g *WeightedGraph) Edges(n *Node) (edges map[Node]int, ok bool) {
	edges, ok = g.edges[*n]
	return
}
