package luhn

import (
	"regexp"
	"strings"
)

// Valid determines whether or not it passes the Luhn checksum formula used to
// validate a variety of ID numbers, e.g. Credit cards and Canadian SINs
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	if len(number) == 1 {
		return false
	} else if match, _ := regexp.MatchString("[^0-9]", number); match {
		return false
	}

	sum := 0
	doDbl := false
	for i := len(number) - 1; i >= 0; i-- {
		d := int(number[i] - '0')

		if doDbl {
			if d = d * 2; d > 9 {
				d -= 9
			}
		}

		sum += d
		doDbl = !doDbl
	}
	return sum%10 == 0
}
