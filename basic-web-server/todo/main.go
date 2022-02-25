package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, PageData{
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Make Template", Done: false},
		},
	})
}

func main() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	fs := http.FileServer(http.Dir("./static/"))
	mux := http.NewServeMux()
	// Necessary because it would come out as ./static/static/css otherwise
	// XXX: This will NOT run without prefix and suffic `/`
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	log.Fatal(http.ListenAndServe(":9922", mux))
}
