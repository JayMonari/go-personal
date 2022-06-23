package bike

import "fmt"

type Bicycle struct {
	Brand, Color string
}

func (b *Bicycle) Move() {
	fmt.Println("Onwards we go!")
}
