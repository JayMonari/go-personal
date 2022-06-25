package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on http://localhot:3000/index.html")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
