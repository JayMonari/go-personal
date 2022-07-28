package raindrops

import "strconv"

func Convert(number int) string {
	sounds := ""
	if number%3 == 0 {
		sounds += "Pling"
	}
	if number%5 == 0 {
		sounds += "Plang"
	}
	if number%7 == 0 {
		sounds += "Plong"
	}

	if sounds == "" {
		return strconv.Itoa(number)
	}
	return sounds
}
