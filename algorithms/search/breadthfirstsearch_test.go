package search

import "testing"

func TestBreadFirstSearch(t *testing.T) {
	you := Node{name: "you", profile: "customer"}
	alice := Node{name: "alice", profile: "designer"}
	bob := Node{name: "bob", profile: "solder"}
	peggy := Node{name: "peggy", profile: "mango_seller"}
	claire := Node{name: "claire", profile: "developer"}
	tom := Node{name: "tom", profile: "student"}

	graph := NewUnweightedGraph()
	graph.AddEdges(&you, []*Node{&alice, &claire, &bob})
	graph.AddEdges(&alice, []*Node{&peggy})
	graph.AddEdges(&claire, []*Node{&tom})
	graph.AddEdges(&tom, []*Node{&peggy})

	mangoSeller, ok := BreadthFirstSearch(
		&you,
		graph,
		func(n *Node) bool {
			return n.profile == "mango_seller"
		},
	)

	if !ok {
		t.Fatalf("Not found: %s", peggy.name)
	}

	if *mangoSeller != peggy {
		t.Fatalf("Wrong result: %s, must be %s", mangoSeller.name, peggy.name)
	}

}

func TestBreadFirstSearchNotFound(t *testing.T) {
	you := Node{name: "you", profile: "customer"}
	alice := Node{name: "alice", profile: "designer"}
	bob := Node{name: "bob", profile: "solder"}
	peggy := Node{name: "peggy", profile: "mango_seller"}
	claire := Node{name: "claire", profile: "developer"}
	tom := Node{name: "tom", profile: "student"}

	graph := NewUnweightedGraph()
	graph.AddEdges(&you, []*Node{&claire, &bob})
	graph.AddEdges(&alice, []*Node{&peggy})
	graph.AddEdges(&claire, []*Node{&tom})

	mangoSeller, ok := BreadthFirstSearch(
		&you,
		graph,
		func(n *Node) bool {
			return n.profile == "mango_seller"
		},
	)

	if ok {
		t.Fatalf("Result must be not found, get: %s", mangoSeller.name)
	}
}
