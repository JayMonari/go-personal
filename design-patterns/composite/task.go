package composite

import (
	"fmt"
	"strings"
)

type Task struct {
	Name     string
	Anchor   string
	Total    int
	SubTasks []Item
}

func (t *Task) Add(i Item) { t.SubTasks = append(t.SubTasks, i) }

func (t *Task) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\t|--%s - %s $%d\n", t.Name, t.Anchor, t.Price()))
	for _, st := range t.SubTasks {
		sb.WriteString(st.String())
	}
	return sb.String()
}

func (t *Task) Price() int {
	pr := t.Total
	for _, st := range t.SubTasks {
		pr += st.Price()
	}
	return pr
}
