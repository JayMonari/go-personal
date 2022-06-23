package todo

import "bridge/list"

type Unique struct {
	rendering list.List
	todos     []string
}

func NewUnique() *Unique { return &Unique{} }

func (u *Unique) SetList(l list.List) { u.rendering = l }

func (u *Unique) Add(td string) {
	for _, t := range u.todos {
		if td == t {
			return
		}
	}
	u.todos = append(u.todos, td)
}

func (u *Unique) Print() { u.rendering.Print(u.todos) }
