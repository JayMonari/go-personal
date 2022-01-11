package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// This package has been deprecated but to understand what might be seen as
// utilities we'll flesh it out here and move it later.
func main() {
  ReadAllExample()
}

// When is this useful? Testing maybe?
func ReadAllExample() {
	r := strings.NewReader("NewReader returns a new Reader reading from s.\nIt is similar to bytes.NewBufferString but more efficient and read-only.")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
