package todo

import "bridge/list"

type Todo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
