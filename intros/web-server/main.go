package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color: steelblue'>Hello</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Book{
		Title:  "The Gunslinger",
		Author: "Stephen King",
		Pages:  304,
	})
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/book", GetBook)

	log.Fatal(http.ListenAndServe(":9922", nil))
}
