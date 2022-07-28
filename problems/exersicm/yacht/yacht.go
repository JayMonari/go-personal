package yacht

// The categories that can be chosen.
const (
	Ones           = "ones"
	Twos           = "twos"
	Threes         = "threes"
	Fours          = "fours"
	Fives          = "fives"
	Sixes          = "sixes"
	FullHouse      = "full house"
	FourOfAKind    = "four of a kind"
	LittleStraight = "little straight"
	BigStraight    = "big straight"
	Choice         = "choice"
	Yacht          = "yacht"
)

// Score returns the score of a given roll based on the rules of Yacht. The
// function panics if an unimplemented category is passed in.
func Score(dice []int, category string) int {
	switch category {
	case Ones:
		return sumN(dice, func(d int) bool { return d == 1 })
	case Twos:
		return sumN(dice, func(d int) bool { return d == 2 })
	case Threes:
		return sumN(dice, func(d int) bool { return d == 3 })
	case Fours:
		return sumN(dice, func(d int) bool { return d == 4 })
	case Fives:
		return sumN(dice, func(d int) bool { return d == 5 })
	case Sixes:
		return sumN(dice, func(d int) bool { return d == 6 })
	case Choice:
		return sumN(dice, func(d int) bool { return true })
	case FullHouse:
		cntr := makeCounter(dice)
		if len(cntr) != 2 {
			return 0
		}

		sum := 0
		for k, v := range cntr {
			if v != 2 && v != 3 {
				return 0
			}
			sum += (k * v)
		}
		return sum

	case FourOfAKind:
		cntr := makeCounter(dice)
		if len(cntr) > 2 {
			return 0
		}

		sum := 0
		for k, v := range cntr {
			if v >= 4 {
				sum = 4 * k
			}
		}
		return sum

	case LittleStraight:
		cntr := makeCounter(dice)
		if _, ok := cntr[6]; ok || len(cntr) != 5 {
			return 0
		}
		return 30

	case BigStraight:
		cntr := makeCounter(dice)
		if _, ok := cntr[1]; ok || len(cntr) != 5 {
			return 0
		}
		return 30

	case Yacht:
		if len(makeCounter(dice)) != 1 {
			return 0
		}
		return 50

	default:
		panic("unimplemented category")
	}
}

// makeCounter creates a counter for the slice of ints.
func makeCounter(dice []int) map[int]int {
	c := make(map[int]int, len(dice))
	for _, d := range dice {
		c[d]++
	}
	return c
}

// sumN sums a slice of ints based on a given predicate.
func sumN(dice []int, pred func(n int) bool) int {
	sum := 0
	for _, d := range dice {
		if pred(d) {
			sum += d
		}
	}
	return sum
}
