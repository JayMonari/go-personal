package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix has methods to return the rows and columns of the matrix along with
// the ability to set values.
type Matrix [][]int

// New makes a matrix from a string of format "1 2 3\n 4 5 6". If the rows are
// not delimited by newlines and the fields by whitespace, then an error will
// be returned. The values must fit into int64 type and the columns must be of
// equal length.
func New(mx string) (*Matrix, error) {
	m := Matrix{}
	for _, line := range strings.Split(mx, "\n") {
		fields := strings.Fields(line)
		if len(m) > 0 && len(fields) != len(m[len(m)-1]) {
			return nil, errors.New("column values do not match")
		}

		row, err := makeRow(fields)
		if err != nil {
			return nil, err
		}

		m = append(m, row)
	}
	return &m, nil
}

// Rows creates a new slice of slices of ints of the matrix rows.
func (m *Matrix) Rows() [][]int {
	rr := make([][]int, len(*m))
	for i, row := range *m {
		rr[i] = make([]int, len(row))
		for j, val := range row {
			rr[i][j] = val
		}
	}
	return rr
}

// Cols creates a new slice of slices of ints of the matrix columns.
func (m *Matrix) Cols() [][]int {
	var nCols int
	if (len(*m)) > 0 {
		nCols = len((*m)[0])
	}
	cc := make([][]int, nCols)
	for _, row := range *m {
		for i, col := range row {
			cc[i] = append(cc[i], col)
		}
	}
	return cc
}

// Set returns true if it sets a value in the matrix in constant time if it is
// within the row and column bounds, false otherwise.
func (m *Matrix) Set(r, c, val int) bool {
	if r < 0 || c < 0 || r >= len(*m) || (len(*m) > 0 && c >= len((*m)[0])) {
		return false
	}
	(*m)[r][c] = val
	return true
}

// makeRow takes a slice of strings as digits and returns them as a slice of
// ints. If the string is not a number the error is returned.
func makeRow(fields []string) ([]int, error) {
	row := make([]int, len(fields))
	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		row[i] = n
	}
	return row, nil
}
