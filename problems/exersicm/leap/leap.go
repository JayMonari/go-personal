// Package leap has IsLeapYear which tells if a given year is a leap year or
// not.
package leap

// IsLeapYear returns true if the year is a leap year, false otherwise.
func IsLeapYear(y int) bool {
	if y%4 != 0 {
		return false
	}
	if y%100 != 0 {
		return true
	}
	if y%400 == 0 {
		return true
	}
	return false
}
