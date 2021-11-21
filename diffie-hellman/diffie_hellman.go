package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var bigTwo = big.NewInt(2)

// PrivateKey returns a random private key in range [2, p).
func PrivateKey(p *big.Int) *big.Int {
	max := new(big.Int).Sub(p, bigTwo)
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return random.Add(bigTwo, random)
}

// PublicKey returns the public key corresponding to the private key a, using:
// g**a mod p
func PublicKey(a, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), a, p)
}

// NewPair generates a new private/public keypair.
func NewPair(p *big.Int, g int64) (priv, pub *big.Int) {
	priv = PrivateKey(p)
	return priv, PublicKey(priv, p, g)
}

// SecretKey returns the public key for two participants using:
// B**a mod p
func SecretKey(a, B, p *big.Int) *big.Int {
	return new(big.Int).Exp(B, a, p)
}
