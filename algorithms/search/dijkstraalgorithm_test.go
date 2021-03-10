package search_test

import (
	"gnotes/algorithms/search"
	"reflect"
	"testing"
)

func TestDijkstraAlgorithm(t *testing.T) {
	you := search.Node{Name: "you", Profile: "customer"}
	alice := search.Node{Name: "alice", Profile: "designer"}
	bob := search.Node{Name: "bob", Profile: "solder"}
	peggy := search.Node{Name: "peggy", Profile: "mango_seller"}
	claire := search.Node{Name: "claire", Profile: "developer"}
	tom := search.Node{Name: "tom", Profile: "student"}

	graph := search.NewWeightedGraph()
	graph.AddEdge(&you, &alice, 2)
	graph.AddEdge(&alice, &peggy, 15)

	graph.AddEdge(&you, &claire, 3)
	graph.AddEdge(&claire, &peggy, 20)

	graph.AddEdge(&you, &bob, 5)
	graph.AddEdge(&bob, &tom, 4)
	graph.AddEdge(&tom, &peggy, 1)

	// final node
	graph.AddEdge(&peggy, nil, 1)

	shortcut, cost := search.DijkstraAlgorithm(
		graph,
		&you,
		&peggy,
	)

	wantWay := []*search.Node{&you, &bob, &tom, &peggy}
	wantCost := 10

	if !reflect.DeepEqual(shortcut, wantWay) {
		t.Fatalf("Wrong shortest way %v, result must be %v", shortcut, wantWay)
	}

	if cost != wantCost {
		t.Fatalf("Wrong cost %v, result must be %v", cost, wantCost)
	}
}
