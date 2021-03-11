package issue

// SetCover greedy algorithm
// O(n^2)
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

// O(a+b)
func intersection(a []string, b []string) []string {
	list := append(a, b...)
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

// O(a+b)
func difference(a []string, b []string) []string {
	d := make([]string, 0)
	m := make(map[string]bool)

	for _, val := range b {
		if _, ok := m[val]; !ok {
			m[val] = true
		}
	}

	for _, val := range a {
		if _, ok := m[val]; !ok {
			d = append(d, val)
		}
	}

	return d
}

// O(n)
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
