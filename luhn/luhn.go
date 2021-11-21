package luhn

import (
	"regexp"
	"strings"
)

// Valid determines wheter or not it passes the Luhn checksum formula used to
// calidate a variety of ID numbers, e.g. Credit cards and Canadian SINs
func Valid(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	if len(number) == 1 {
		return false
	} else if match, _ := regexp.MatchString("[^0-9]", number); match {
		return false
	}

	sum := 0
	double := false
	for i := len(number) - 1; i >= 0; i-- {
		if double {
			doubled := int(number[i]-'0') * 2
			if doubled > 9 {
				doubled -= 9
			}
			sum += doubled
		} else {
			sum += int(number[i] - '0')
		}
		double = !double
	}
	return sum%10 == 0
}
