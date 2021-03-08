package search

import "sync"

// BreadthFirstSearch realization
func BreadthFirstSearch(from *Node, graph *Graph, handler func(*Node) bool) (*Node, bool) {
	edges, ok := graph.edges[*from]
	if !ok {
		return nil, false
	}

	nodeQueue := NewNodeQueue()
	nodeQueue.Enqueue(edges...)
	searched := make(map[Node]bool)

	for !nodeQueue.IsEmpty() {
		node := nodeQueue.Dequeue()
		if _, ok := searched[*node]; ok {
			continue
		}

		if handler(node) {
			return node, true
		}

		searched[*node] = true

		edges, ok = graph.edges[*node]
		if ok {
			nodeQueue.Enqueue(edges...)
		}
	}

	return nil, false
}

// NodeQueue the queue of Nodes
type NodeQueue struct {
	nodes []*Node
	lock  sync.RWMutex
}

// NewNodeQueue creates a new NodeQueue
func NewNodeQueue() *NodeQueue {
	return &NodeQueue{
		nodes: make([]*Node, 0),
		lock:  sync.RWMutex{},
	}
}

// Enqueue adds an Node to the end of the queue
func (q *NodeQueue) Enqueue(t ...*Node) {
	q.lock.Lock()
	q.nodes = append(q.nodes, t...)
	q.lock.Unlock()
}

// Dequeue removes an Node from the start of the queue
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	item := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.lock.Unlock()

	return item
}

// Front returns the item next in the queue, without removing it
func (q *NodeQueue) Front() *Node {
	q.lock.RLock()
	item := q.nodes[0]
	q.lock.RUnlock()

	return item
}

// IsEmpty returns true if the queue is empty
func (q *NodeQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.nodes) == 0
}

// Size returns the number of Nodes in the queue
func (q *NodeQueue) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.nodes)
}
