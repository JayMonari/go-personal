package ecs

import "reflect"

type ID uint32

type Component interface {
	ComponentSet(any)
}

type BasicStorage struct {
	list map[ID]any
}

func NewBasicStorage() *BasicStorage {
	return &BasicStorage{
		list: make(map[ID]any),
	}
}

func (s *BasicStorage) Read(id ID) (any, bool) {
	val, ok := s.list[id]
	return val, ok
}

func (s *BasicStorage) Write(id ID, val any) { s.list[id] = val }

type Engine struct {
	req       map[string]*BasicStorage
	idCounter ID
}

func NewEngine() *Engine {
	return &Engine{
		req: make(map[string]*BasicStorage),
	}
}

func (e *Engine) NewID() ID {
	id := e.idCounter
	e.idCounter++
	return id
}

func name(t any) string {
	name := reflect.TypeOf(t).String()
	if name[0] == '*' {
		return name[1:]
	}
	return name
}

func GetStorage(e *Engine, t any) *BasicStorage {
	name := name(t)
	s, ok := e.req[name]
	if !ok {
		e.req[name] = NewBasicStorage()
		s = e.req[name]
	}
	return s
}

func Read(e *Engine, id ID, val Component) bool {
	newVal, ok := GetStorage(e, val).Read(id)
	if ok {
		val.ComponentSet(newVal)
	}
	return ok
}

func Write(e *Engine, id ID, val any) { GetStorage(e, val).Write(id, val) }

func Each(e *Engine, t any, f func(id ID, a any)) {
	for id, a := range GetStorage(e, t).list {
		f(id, a)
	}
}
