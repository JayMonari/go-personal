package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Garden is a map of a child and the plants they are responsible for in the
// kindergarten garden.
type Garden map[string][]string

// Plants returns the plants owned by the child from the garden.
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}

// plantNames holds the names of each plant mapped by the first Uppercase rune.
var plantNames = map[rune]string{
	'C': "clover",
	'G': "grass",
	'R': "radishes",
	'V': "violets",
}

// NewGarden creates a Garden out of a provided diagram of the two rows of the
// kindergarten garden. The children will be mapped to each of their respective
// plants in the order provided. If the diagram format is incorrect, rows are
// not equal or of odd length, or children has duplicate names an error is
// returned.
func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, fmt.Errorf("invalid garden format")
	}

	rows, err := split(diagram)
	if err != nil {
		return nil, err
	}

	sortedKids, err := order(children)
	if err != nil {
		return nil, err
	}

	garden := Garden{}
	for i := 0; i < 2; i++ {
		if 2*len(children) != len(rows[i]) {
			return nil, fmt.Errorf("each child must have 2 plants per rows")
		}
		if err := assign(sortedKids, rows[i], garden); err != nil {
			return nil, err
		}
	}
	return &garden, nil
}

// assign fills the Garden with the seeds of the row for each child. An error
// is returned if an invalid seed was found in row.
func assign(kids []string, row string, g Garden) error {
	for j, seed := range row {
		name, ok := plantNames[seed]
		if !ok {
			return fmt.Errorf("unknown seed %b passed in", seed)
		}
		kid := kids[j/2]
		plantList, ok := g[kid]
		if !ok {
			plantList = []string{}
		}
		g[kid] = append(plantList, name)
	}
	return nil
}

// order sorts the children and returns a new slice of strings. An error is
// returned if duplicate names are found in the slice.
func order(children []string) ([]string, error) {
	srt := make([]string, len(children))
	copy(srt, children)
	sort.Strings(srt)
	if hasDuplicate(srt) {
		return nil, fmt.Errorf("")
	}
	return srt, nil
}

// split returns the diagram split into two rows. If the amount of rows does
// not equal two or the length of both rows is not equal an error is returned.
func split(diagram string) ([]string, error) {
	rows := strings.Split(diagram[1:], "\n")
	switch {
	case len(rows) != 2:
		return nil, fmt.Errorf("garden needs two rows %d rows found", len(rows))
	case len(rows[0]) != len(rows[1]):
		return nil, fmt.Errorf("rows uneven %d and %d", len(rows[0]), len(rows[1]))
	}
	return rows, nil
}

// hasDuplicate checks whether or not a sorted slice of string has a duplicate
// string.
func hasDuplicate(sl []string) bool {
	for i := 0; i < len(sl)-1; i++ {
		if sl[i] == sl[i+1] {
			return true
		}
	}
	return false
}
