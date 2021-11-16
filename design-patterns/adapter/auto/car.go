package auto

import "fmt"

type Car struct {
	Brand   string
	Model   uint8
	Started bool
}

func (c *Car) Start() {
	if c.Started {
		fmt.Println("The car is already started!")
		return
	}
	c.Started = true
	fmt.Println("Car started!")
}

func (c *Car) Accel() { fmt.Println("Vroom Vroom") }
