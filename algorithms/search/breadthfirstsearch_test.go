package search_test

import (
	"gnotes/algorithms/search"
	"testing"
)

func TestBreadFirstSearch(t *testing.T) {
	you := search.Node{Name: "you", Profile: "customer"}
	alice := search.Node{Name: "alice", Profile: "designer"}
	bob := search.Node{Name: "bob", Profile: "solder"}
	peggy := search.Node{Name: "peggy", Profile: "mango_seller"}
	claire := search.Node{Name: "claire", Profile: "developer"}
	tom := search.Node{Name: "tom", Profile: "student"}

	graph := search.NewUnweightedGraph()
	graph.AddEdges(&you, []*search.Node{&alice, &claire, &bob})
	graph.AddEdges(&alice, []*search.Node{&peggy})
	graph.AddEdges(&claire, []*search.Node{&tom})
	graph.AddEdges(&tom, []*search.Node{&peggy})

	mangoSeller, ok := search.BreadthFirstSearch(
		&you,
		graph,
		func(n *search.Node) bool {
			return n.Profile == "mango_seller"
		},
	)

	if !ok {
		t.Fatalf("Not found: %s", peggy.Name)
	}

	if *mangoSeller != peggy {
		t.Fatalf("Wrong result: %s, must be %s", mangoSeller.Name, peggy.Name)
	}

}

func TestBreadFirstSearchNotFound(t *testing.T) {
	you := search.Node{Name: "you", Profile: "customer"}
	alice := search.Node{Name: "alice", Profile: "designer"}
	bob := search.Node{Name: "bob", Profile: "solder"}
	peggy := search.Node{Name: "peggy", Profile: "mango_seller"}
	claire := search.Node{Name: "claire", Profile: "developer"}
	tom := search.Node{Name: "tom", Profile: "student"}

	graph := search.NewUnweightedGraph()
	graph.AddEdges(&you, []*search.Node{&claire, &bob})
	graph.AddEdges(&alice, []*search.Node{&peggy})
	graph.AddEdges(&claire, []*search.Node{&tom})

	mangoSeller, ok := search.BreadthFirstSearch(
		&you,
		graph,
		func(n *search.Node) bool {
			return n.Profile == "mango_seller"
		},
	)

	if ok {
		t.Fatalf("Result must be not found, get: %s", mangoSeller.Name)
	}
}
