package decorator

import "fmt"

type HandlerBye struct{}

func (h *HandlerBye) Process() error {
  fmt.Println("good bai >:)")
	return nil
}
