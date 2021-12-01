package main

import (
	"adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA app")
	name := flag.String("name", "gopher.json", "the JSON name with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *name)

	f, err := os.Open(*name)
	if err != nil {
		panic(err)
	}

	story, err := adventure.JSONStory(f)
	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story)
	// These would need the story to be mapped in every chapter
	// adventure.WithPathFunc(pathFn))
	// mux := http.NewServeMux()
	// mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

// pathFn dynamically change path of JSON file endpoints; won't work in its
// current condition.
func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
