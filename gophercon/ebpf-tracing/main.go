package main

import (
	"fmt"
	"log"
	"net/http"
)

//go:noinline
func handlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Recieved connection from %s\n", r.Host)
}

// In order to make this work you will need both `strace` and `bcc` installed
func main() {
	http.HandleFunc("/", handlerFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
