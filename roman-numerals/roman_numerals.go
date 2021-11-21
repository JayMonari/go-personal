package romannumerals

import (
	"errors"
	"strings"
)

// NumeralConversion holds the letter and value of a roman numeral.
type NumeralConversion struct {
	roman  string
	arabic int
}

var numerals = []NumeralConversion{
	{roman: "M", arabic: 1000},
	{roman: "CM", arabic: 900},
	{roman: "D", arabic: 500},
	{roman: "CD", arabic: 400},
	{roman: "C", arabic: 100},
	{roman: "XC", arabic: 90},
	{roman: "L", arabic: 50},
	{roman: "XL", arabic: 40},
	{roman: "X", arabic: 10},
	{roman: "IX", arabic: 9},
	{roman: "V", arabic: 5},
	{roman: "IV", arabic: 4},
	{roman: "I", arabic: 1},
}

// ToRomanNumeral takes in an integer from 1 to 3000 and gives back a string as
// a roman numeral. If the value is out of range an error is returned.
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3000 {
		return "", errors.New("number must be between 1-3000.")
	}

	numeral := strings.Builder{}
	for _, conv := range numerals {
		for num >= conv.arabic {
			numeral.WriteString(conv.roman)
			num -= conv.arabic
		}
	}

	return numeral.String(), nil
}
