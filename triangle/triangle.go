package triangle

import "math"

// Kind creates the Enum for the different kinds of triangles
type Kind uint8

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides takes in the three sides of a triangle and returns what kind
// of triangle it is. If the three sides do not make a triangle then the value
// for Not a Triangle is returned.
func KindFromSides(a, b, c float64) Kind {
	if isInequal(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func isInequal(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return true
	} else if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return true
	} else if a == 0 || b == 0 || c == 0 {
		return true
	} else if a+b < c || b+c < a || a+c < b {
		return true
	}
	return false
}
