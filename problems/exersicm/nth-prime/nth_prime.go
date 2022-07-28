package prime

import (
	"math"
)

// limit is the highest a prime number can reach before it is out of bounds.
const limit = 105000

// Nth returns the nth prime number.
func Nth(n int) (int, bool) {
	if n < 1 {
		return -1, false
	}

	isPrime := sieveOfAtkin(n + 1)
	primes := make([]int, 0, n)
	for p := 0; p < len(isPrime)-1; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes[n-1], true
}

// sieveOfAtkin generates an array of prime numbers up to limit.
func sieveOfAtkin(m int) [limit]bool {
	isPrime := [limit]bool{}
	isPrime[2] = true
	isPrime[3] = true

	nsqrt := math.Sqrt(limit)
	for x := 1; float64(x) <= nsqrt; x++ {
		for y := 1; float64(y) <= nsqrt; y++ {
			n := 4*(x*x) + y*y
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= limit && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= limit && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}
	// Remove all factors as they are not primes.
	for n := 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y := n * n; y < limit; y += n * n {
				isPrime[y] = false
			}
		}
	}
	return isPrime
}
