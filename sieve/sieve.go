package sieve

// Sieve returns all of the prime numbers up to and including n.
func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	primes := make([]int, 0, n/2)
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			continue
		}
		primes = append(primes, i)
		for fac := i; fac <= n; fac += i {
			isPrime[fac] = true
		}
	}
	return primes
}
