package darts

// Score returns the amount of points rewarded to player who threw the dart,
// depending on where it lands on the board.
func Score(x, y float64) int {
	point := x*x + y*y
	switch {
	case point <= 1:
		return 10
	case point <= 25:
		return 5
	case point <= 100:
		return 1
	default:
		return 0
	}
}
