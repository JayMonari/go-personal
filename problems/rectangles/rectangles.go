package rectangles

// Count returns the number of rectangles found in a given diagram.
func Count(diagram []string) (count int) {
	for i, line := range diagram {
		var points []int
		for j, r := range line {
			if r == '+' {
				points = append(points, j)
			}
		}

		for p1 := 0; p1 < len(points)-1; p1++ {
			for p2 := p1 + 1; p2 < len(points); p2++ {
				for j := i + 1; j < len(diagram); j++ {
					j1, j2 := diagram[j][points[p1]], diagram[j][points[p2]]
					if j1 == '+' && j2 == '+' {
						count++
					} else if (j1 == '+' && j2 == '|') || (j1 == '|' && j2 == '+') {
						continue
					} else if j1 != '|' || j2 != '|' {
						break
					}
				}
			}
		}
	}
	return count
}
