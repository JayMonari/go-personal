package adapter

import (
	"adapter/auto"
	"adapter/bike"
)

// Factory returns an adapter for a vehicle if it exists, nil otherwise.
func Factory(s string) Adapter {
	switch s {
	case "car":
		d := auto.Car{}
		return &CarAdapter{&d}
	case "bike":
		d := bike.Bicycle{}
		return &BikeAdapter{&d}
	}
	return nil
}
