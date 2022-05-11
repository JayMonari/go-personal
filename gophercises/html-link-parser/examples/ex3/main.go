package main

import (
	"fmt"
	"link"
	"os"
)

func main() {
	f, err := os.Open("./examples/ex3/ex3.html")
	if err != nil {
		panic(err)
	}
	ll, err := link.Parse(f)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("%+v", ll)
}
