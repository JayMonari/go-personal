package securerandom

import (
	"crypto/rand"
	"math/big"
)

// Intn is a shortcut for generating a random integer between 0 and max using
// crypto/rand.
func Intn(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

// Float64 is a shortcut for generating a random float between 0 and 1 using
// crypto/rand.
func Float64() float64 { return float64(Intn(1<<53) / (1 << 53)) }
