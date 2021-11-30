package main

import (
	"adventure"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
