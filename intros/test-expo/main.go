package main

import (
	"testexpo/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handler.Add)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
