package spiralmatrix

// SpiralMatrix returns a matrix of natural numbers that starts at 1 at the top
// left corner and spirals inward going from left to right and down to up. The
// size given is the amount of rows the matrix has.
func SpiralMatrix(size int) [][]int {
	m := make([][]int, size)
	for row := 0; row < size; row++ {
		m[row] = make([]int, size)
	}
	row, col := 0, 0
	dr, dc := 0, 1
	for n := 1; n <= size*size; n++ {
		m[row][col] = n
		if outOfBounds(size, row, dr, col, dc) || m[row+dr][col+dc] != 0 {
			dr, dc = dc, -dr // change direction
		}
		row, col = row+dr, col+dc
	}
	return m
}

func outOfBounds(size, r, dr, c, dc int) bool {
	return 0 > r+dr || r+dr >= size || 0 > c+dc || c+dc >= size
}
