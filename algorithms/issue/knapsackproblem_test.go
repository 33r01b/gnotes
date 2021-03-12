package issue_test

import (
	"gnotes/algorithms/issue"
	"reflect"
	"sort"
	"testing"
)

func TestKnapsackProblem(t *testing.T) {
	player := issue.Item{
		Name:   "player",
		Cost:   3000,
		Weight: 4,
	}
	notebook := issue.Item{
		Name:   "notebook",
		Cost:   2000,
		Weight: 3,
	}
	guitar := issue.Item{
		Name:   "guitar",
		Cost:   1500,
		Weight: 1,
	}
	phone := issue.Item{
		Name:   "phone",
		Cost:   2000,
		Weight: 1,
	}

	items := []issue.Item{player, notebook, guitar, phone}
	knapsack := issue.Knapsack{Capacity: 4, Items: make([]issue.Item, 0)}

	issue.KnapsackProblem(&knapsack, items)
	sort.Slice(knapsack.Items, func(i, j int) bool {
		return knapsack.Items[i].Weight < knapsack.Items[j].Weight
	})

	wantItems := []issue.Item{notebook, phone}
	sort.Slice(wantItems, func(i, j int) bool {
		return wantItems[i].Weight < wantItems[j].Weight
	})

	if !reflect.DeepEqual(knapsack.Items, wantItems) {
		t.Fatalf("Wrong result %v, items must be %v", knapsack.Items, wantItems)
	}

	wantSum := 4000

	if knapsack.Sum != wantSum {
		t.Fatalf("Wrong result %v, sum must be %v", knapsack.Sum, wantSum)
	}
}
