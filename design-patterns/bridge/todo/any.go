package todo

import "bridge/list"

type Any struct {
	rendering list.List
	todos     []string
}

func NewAny() *Any {
	return &Any{}
}

func (a *Any) SetList(l list.List) { a.rendering = l }

func (a *Any) Add(td string) { a.todos = append(a.todos, td) }

func (a *Any) Print() { a.rendering.Print(a.todos) }
