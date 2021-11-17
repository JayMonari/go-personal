package todo

import "bridge/list"

type Important struct {
	rendering list.List
	todos     []string
}

func NewImportant() *Important { return &Important{} }

func (im *Important) SetList(l list.List) { im.rendering = l }

func (im *Important) Add(td string) {
	for i, t := range im.todos {
		if td == t {
			im.todos[i] = "**VERY IMPORTANT**" + td
			return
		}
	}
	im.todos = append(im.todos, td)
}

func (im *Important) Print() { im.rendering.Print(im.todos) }
