package todo

import (
	"bridge/list"
	"fmt"
	"time"
)

type Dated struct {
	rendering list.List
	todos     []string
	format    string
}

func NewDated(format string) *Dated {
	if format == "" {
		format = "Monday, January 2, 3:04PM"
	}
	return &Dated{format: format}
}

func (d *Dated) SetList(l list.List) { d.rendering = l }

func (d *Dated) Add(td string) {
	date := time.Now().Local().Format(d.format)
	d.todos = append(d.todos, fmt.Sprintf("%s %s", date, td))
}

func (d *Dated) Print() { d.rendering.Print(d.todos) }
