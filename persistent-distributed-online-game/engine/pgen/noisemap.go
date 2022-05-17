package pgen

import (
	"github.com/ojrac/opensimplex-go"
)

type NoiseMap struct {
	seed     int64
	noise    opensimplex.Noise
	exponent float64
}

func NewNoiseMap(seed int64, exponent float64) *NoiseMap {
	return &NoiseMap{
		seed:     seed,
		noise:    opensimplex.NewNormalized(seed),
		exponent: exponent,
	}
}

func (n *NoiseMap) Get(x, y int) float64 {
	freq := 0.01
	return n.noise.Eval2(freq*float64(x), freq*float64(y))
}
