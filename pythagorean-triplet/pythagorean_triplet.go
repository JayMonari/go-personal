package pythagorean

import (
	"math"
	"sort"
)

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the range min to max inclusive.
func Range(min, max int) (triplets []Triplet) {
	for m := intSqrt(min / 2); m <= intSqrt(max/2); m++ {
		for n := 1 + (m % 2); n < m; n += 2 {
			if gcd(m, n) != 1 {
				continue
			}
			for k := 1; k <= (max / (m*m + n*n)); k++ {
				a, b, c := euclid(k, m, n)
				if a >= min && b >= min {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	if p%2 == 1 {
		return nil
	}
	p /= 2
	divs := divisors(p)
	sort.Ints(divs)
	var triplets []Triplet
	for i, m := range divs {
		if m > intSqrt(p) {
			break
		}
		for j := i + 1; j < len(divs); j++ {
			if divs[j] >= 2*m {
				break
			}
			if p%(m*divs[j]) == 0 {
				n := divs[j] - m
				k := p / m / divs[j]
				a, b, c := euclid(k, m, n)
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return removeDups(triplets)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func intSqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func euclid(k, m, n int) (a, b, c int) {
	a = k * 2 * m * n
	b = k * (m*m - n*n)
	c = k * (m*m + n*n)
	if a > b {
		a, b = b, a
	}
	return
}

func divisors(p int) (divs []int) {
	for i := 1; i <= intSqrt(p); i++ {
		if p%i == 0 {
			if p/i == i {
				divs = append(divs, i)
			} else {
				divs = append(divs, i, p/i)
			}
		}
	}
	return divs
}
func removeDups(triplets []Triplet) []Triplet {
	sort.Slice(triplets, func(i, j int) bool { return triplets[i][0] < triplets[j][0] })
	result := triplets[:1]
	for i := 1; i < len(triplets); i++ {
		if triplets[i] != triplets[i-1] {
			result = append(result, triplets[i])
		}
	}
	return result
}
