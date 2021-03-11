package issue

// SetCover greedy algorithm
func SetCover(statesNeed []string, stations map[string][]string) []string {
	finalStations := make([]string, 0, len(statesNeed))

	for len(statesNeed) > 0 {
		bestStation := ""
		statesCovered := make([]string, 0)

		for station, statesForStation := range stations {
			// find intersections
			covered := intersection(statesNeed, statesForStation)

			// select biggest
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		// keep only differences
		statesNeed = difference(statesNeed, statesCovered)
		// add best
		finalStations = append(finalStations, bestStation)
	}

	return unique(finalStations)
}

func intersection(foo []string, bar []string) []string {
	list := append(foo, bar...)
	i := make([]string, 0)
	m := make(map[string]bool)

	for _, val := range list {
		if _, ok := m[val]; !ok {
			m[val] = true
		} else {
			i = append(i, val)
		}
	}

	return i
}

func difference(foo []string, bar []string) []string {
	d := make([]string, 0)
	m := make(map[string]bool)

	for _, val := range bar {
		if _, ok := m[val]; !ok {
			m[val] = true
		}
	}

	for _, val := range foo {
		if _, ok := m[val]; !ok {
			d = append(d, val)
		}
	}

	return d

}

func unique(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}
