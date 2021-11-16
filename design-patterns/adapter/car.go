package adapter

import "adapter/auto"

type CarAdapter struct{ car *auto.Car }

func (c *CarAdapter) Move() {
	c.car.Start()
	c.car.Accel()
}
