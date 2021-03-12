package issue

import "math"

type Knapsack struct {
	Capacity int
	Sum      int
	Items    []Item
}

type Item struct {
	Name   string
	Cost   int
	Weight int
}

// KnapsackProblem dynamic algorithm realization
func KnapsackProblem(knapsack *Knapsack, items []Item) {
	// item -> capacity -> cost
	table := map[int]map[int]Knapsack{}
	minWeight := findMinWeight(items)

	for i, item := range items {
		for j := minWeight; j <= knapsack.Capacity; j += minWeight {
			if _, ok := table[i]; !ok {
				table[i] = map[int]Knapsack{}
				table[i][j] = Knapsack{
					Capacity: 0,
					Sum:      0,
					Items:    make([]Item, 0),
				}
			}

			// filter oversize items
			if item.Weight > j {
				continue
			}

			var prevCost int
			var prevItems []Item
			cost := item.Cost
			packItems := []Item{item}

			// get previous item
			if _, ok := table[i-1]; ok {
				prevCost = table[i-1][j].Sum
				prevItems = table[i-1][j].Items
			}

			if r, ok := table[i-1][j-item.Weight]; ok {
				cost += r.Sum
				packItems = append(packItems, r.Items...)
			}

			if prevCost > cost {
				cost = prevCost
				packItems = prevItems
			}

			subPack := table[i][j]
			subPack.Sum = cost
			subPack.Items = packItems

			knapsack.Sum = cost
			knapsack.Items = packItems

			table[i][j] = subPack
		}
	}

	return
}

func findMinWeight(items []Item) int {
	min := math.MaxInt64

	for _, i := range items {
		if i.Weight < min {
			min = i.Weight
		}
	}

	return min
}
