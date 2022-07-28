package chance

import (
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	return rnd.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return rnd.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	aa := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rnd.Shuffle(8, func(i, j int) {
		aa[i], aa[j] = aa[j], aa[i]
	})
	return aa
}
