package prime

// Factors returns all prime factors of n.
func Factors(n int64) []int64 {
	fac := []int64{}
	var f int64 = 2
	for n > 1 {
		if n%f == 0 {
			n /= f
			fac = append(fac, f)
		} else {
			f++
		}
	}
	return fac
}
