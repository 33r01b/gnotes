package search

import (
	"math"
)

// DijkstraAlgorithm realization
func DijkstraAlgorithm(graph *WeightedGraph, from *Node, to *Node) (shortcut []*Node, cost int) {
	costs := make(map[Node]int)
	parents := make(map[Node]*Node)

	// find neighbor nodes
	if edge, ok := graph.Edges(from); ok {
		for n, weight := range edge {
			costs[n] = weight
			parents[n] = from
		}
	}

	// add finish
	costs[*to] = math.MaxInt64
	parents[*to] = nil

	processed := make(map[Node]bool)
	node := findLowestCostNode(costs, processed)

	for node != nil {
		neighbors, ok := graph.Edges(node)
		if !ok {
			break
		}

		nodeCost := costs[*node]

		for n, weight := range neighbors {
			newCost := nodeCost + weight

			// add new nodes
			if _, ok := costs[n]; !ok {
				costs[n] = newCost
				parents[n] = node
			}

			// update by smallest cost
			if newCost < costs[n] {
				costs[n] = newCost
				parents[n] = node
			}
		}

		processed[*node] = true
		node = findLowestCostNode(costs, processed)
	}

	// build shortcut way list
	child := *to
	parent, ok := parents[child]

	for ok {
		c := child
		shortcut = append([]*Node{&c}, shortcut...)
		child = *parent
		parent, ok = parents[child]
	}

	shortcut = append([]*Node{from}, shortcut...)
	cost = costs[*to]

	return
}

func findLowestCostNode(costs map[Node]int, processed map[Node]bool) (lowestNode *Node) {
	lowest := math.MaxInt64

	for n, weight := range costs {
		if _, isProcessed := processed[n]; weight < lowest && !isProcessed {
			lowest = weight
			lowestNode = &n
		}
	}

	return
}
