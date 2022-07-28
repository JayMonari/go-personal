package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (min Product, max Product, err error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if !isPalindrome(p) {
				continue
			}

			switch {
			case p > max.Product:
				max.Product = p
				max.Factorizations = [][2]int{{i, j}}
			case p == max.Product:
				max.Factorizations = append(max.Factorizations, [2]int{i, j})
			}

			if min.Product == 0 {
				min.Product = p
			}
			switch {
			case p < min.Product:
				min.Product = p
				min.Factorizations = [][2]int{{i, j}}
			case p == min.Product:
				min.Factorizations = append(min.Factorizations, [2]int{i, j})
			}
		}
	}
	if min.Product == 0 && max.Product == 0 {
		return Product{}, Product{}, errors.New("no palindromes found")
	}
	return
}

func isPalindrome(n int) bool {
	if n < 0 {
		n = -n
	}
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
