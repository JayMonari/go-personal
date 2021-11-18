package composite

import (
	"fmt"
	"strings"
)

type Project struct {
	Name  string
	Tasks []Item
}

func (p *Project) Add(i Item) { p.Tasks = append(p.Tasks, i) }

func (p *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s $%d\n", p.Name, p.Price()))
	for _, t := range p.Tasks {
		sb.WriteString(t.String())
	}
	return sb.String()
}

func (p *Project) Price() int {
	pr := 0
	for _, t := range p.Tasks {
		pr += t.Price()
	}
	return pr
}
