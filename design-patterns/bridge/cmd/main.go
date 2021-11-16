package main

import (
	"bridge/list"
	"bridge/todo"
)

func main() {
  myTodos := factoryTodo("plain")
  myTodos.SetList(factoryList("numbered"))

  myTodos.Add("Never forget just how important Todos are")
  myTodos.Add("Create another implementation")
  myTodos.Add("Create another representation")
  myTodos.Add("Create tests that show this works")
  myTodos.Add("Understand the Bridge pattern better")
  myTodos.Add("Create another representation")
  myTodos.Add("Show off what's in the README.md")
  myTodos.Print()
}

// factoryTodo contains the implementations
func factoryTodo(s string) todo.Todo {
	if s == "unique" {
		return todo.NewUnique()
	}
	return todo.NewAny()
}

// factoryList contains the representations
func factoryList(s string) list.List {
  switch s {
  case "plain":
    return list.NewPlain()
  case "numbered":
    return list.NewNumbered(1)
  }
	return list.NewBullet('*')
}
