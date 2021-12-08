package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int
type Pair [2]int

func New(mx string) (*Matrix, error) {
	m := Matrix{}
	for _, line := range strings.Split(mx, "\n") {
		fields := strings.Fields(line)
		if len(m) > 0 && len(fields) != len(m[len(m)-1]) {
			return nil, errors.New("irregular matrix")
		}

		row, err := makeRow(fields)
		if err != nil {
			return nil, err
		}
		m = append(m, row)
	}
	return &m, nil
}

func (m *Matrix) rowMaxes() []int {
	maxes := make([]int, len(*m))
	for i, row := range *m {
		max := row[0]
		for _, point := range row {
			if max < point {
				max = point
			}
		}
		maxes[i] = max
	}
	return maxes
}

func (m *Matrix) colMins() []int {
	var nCols int
	if len(*m) > 0 {
		nCols = len((*m)[0])
	}

	cols := make([][]int, nCols)
	for i := 0; i < nCols; i++ {
		for _, row := range *m {
			cols[i] = append(cols[i], row[i])
		}
	}

	mins := make([]int, nCols)
	for i, col := range cols {
		min := col[0]
		for _, point := range col[1:] {
			if min > point {
				min = point
			}
		}
		mins[i] = min
	}
	return mins
}

func (m *Matrix) Saddle() []Pair {
	rm := m.rowMaxes()
	cm := m.colMins()
	var points []Pair
	for i := range *m {
		for j := range (*m)[0] {
			cand := (*m)[i][j]
			if cand == rm[i] && cand == cm[j] {
				points = append(points, Pair{i, j})
			}
		}
	}
	return points
}

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
