package composite

import (
	"fmt"
	"strings"
)

type SubTask struct {
	Name  string
	Total int
}

func (s *SubTask) Add(i Item) { fmt.Println("unable to add more items") }

func (s *SubTask) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\t\t|-- %s\n", s.Name))
	return sb.String()
}

func (s *SubTask) Price() int { return s.Total }
