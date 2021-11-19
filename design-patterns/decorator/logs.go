package decorator

import "fmt"

type LogRegistry struct {
	Handler Decorator
}

func NewLogRegistry(d Decorator) *LogRegistry {
	return &LogRegistry{Handler: d}
}

func (lr *LogRegistry) Process() error {
	fmt.Println("You can add functionality before a handler.")
	defer fmt.Println("And after the handler with a decorator!")
	return lr.Handler.Process()
}
