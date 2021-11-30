package main

import (
	"adventure"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "gopher.json", "the JSON name with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *name)

	f, err := os.Open(*name)
	if err != nil {
		panic(err)
	}

  d := json.NewDecoder(f)
  var story adventure.Story
  if err := d.Decode(&story); err != nil {
    panic(err)
  }

  fmt.Printf("%+v\n", story)
}
