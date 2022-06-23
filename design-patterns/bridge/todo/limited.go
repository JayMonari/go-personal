package todo

import (
	"bridge/list"
)

type Limited struct {
	rendering list.List
	todos     []string
	limit     int
}

func NewLimited(lmt int) *Limited {
	return &Limited{limit: lmt}
}

func (lm *Limited) SetList(l list.List) { lm.rendering = l }

func (lm *Limited) Add(td string) {
	if len(lm.todos) >= lm.limit {
		return
	}
	lm.todos = append(lm.todos, td)
}

func (lm *Limited) Print() { lm.rendering.Print(lm.todos) }
