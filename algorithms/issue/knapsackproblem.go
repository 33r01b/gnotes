package issue

import "math"

type Pack struct {
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
func KnapsackProblem(knapsack *Pack, items []Item) {
	// item -> capacity -> Pack
	table := map[int]map[int]Pack{}
	minWeight := findMinWeight(items)

	for i, item := range items {
		if _, ok := table[i]; !ok {
			table[i] = map[int]Pack{}
		}

		for j := minWeight; j <= knapsack.Capacity; j += minWeight {
			if _, ok := table[i][j]; !ok {
				table[i][j] = Pack{}
			}

			// filter oversize items
			if item.Weight > j {
				continue
			}

			// fill pack
			pack := table[i][j]
			pack.Sum = item.Cost
			pack.Items = []Item{item}

			// get previous pack
			prevPack := Pack{}
			if _, ok := table[i-1]; ok {
				prevPack = table[i-1][j]
			}

			// fill with leftovers
			if lo, ok := table[i-1][j-item.Weight]; ok {
				pack.Sum += lo.Sum
				pack.Items = append(pack.Items, lo.Items...)
			}

			// if previous biggest, get them
			if prevPack.Sum > pack.Sum {
				pack.Sum = prevPack.Sum
				pack.Items = prevPack.Items
			}

			table[i][j] = pack

			// fill knapsack, last iteration result must be optimal
			knapsack.Sum = pack.Sum
			knapsack.Items = pack.Items
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
