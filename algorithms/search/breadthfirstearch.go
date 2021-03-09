package search

// BreadthFirstSearch realization
func BreadthFirstSearch(from *Node, graph *UnweightedGraph, handler func(*Node) bool) (*Node, bool) {
	edges, ok := graph.Edges(from)
	if !ok {
		return nil, false
	}

	searched := make(map[Node]bool)

	for len(edges) > 0 {
		node := edges[0]
		edges = edges[1:]

		if _, ok := searched[*node]; ok {
			continue
		}

		if handler(node) {
			return node, true
		}

		searched[*node] = true

		if nodeEdges, ok := graph.Edges(node); ok {
			edges = append(edges, nodeEdges...)
		}
	}

	return nil, false
}
