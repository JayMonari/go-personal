package decorator

import "fmt"

type HandlerHello struct{}

func (h *HandlerHello) Process() error {
	fmt.Println("Oh Hai, Mark")
	return nil
}
