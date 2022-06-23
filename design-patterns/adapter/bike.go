package adapter

import "adapter/bike"

type BikeAdapter struct {
	Bike *bike.Bicycle
}

func (b *BikeAdapter) Move() { b.Bike.Move() }
