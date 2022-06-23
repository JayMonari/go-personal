package main

import (
	"decorator"
	"fmt"
	"log"
)

func main() {
	route := decorator.NewRoute()
	start(&route)

	var path string
	_, err := fmt.Scan(&path)
	if err != nil {
		log.Fatal(err)
	}
	route.Exec(path)
}

func start(route *decorator.Route) {
	route.Add(decorator.NewLogRegistry(&decorator.HandlerHello{}), "/hello")
	route.Add(decorator.NewLogRegistry(&decorator.HandlerBye{}), "/bye")
}
