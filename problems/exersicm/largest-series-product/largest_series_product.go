package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct given a string of digits 0-9 and a span. The max value
// from the products produced from each series is returned.
//
// e.g. digits = "293" span = 2 returns the max of [ 2*9, 9*3 ], which is 27.
func LargestSeriesProduct(number string, span int) (int, error) {
	if span < 0 || len(number) < span {
		return 0, errors.New("span is out of range")
	}

	digits, err := convertToSlice(number)
	if err != nil {
		return 0, err
	}

	products := make([]int, 0, len(digits)-span)
	for i := 0; i <= len(digits)-span; i++ {
		product := multipy(digits[i : i+span])
		products = append(products, product)
	}

	return max(products), nil
}

// convertToSlice converts a given string of digits into an int slice. If the
// string contains anything other than digits 0-9 an error is returned
func convertToSlice(num string) ([]int, error) {
	digits := make([]int, 0, len(num))
	for _, d := range num {
		if !unicode.IsDigit(d) {
			return nil, errors.New("number must only have digits")
		}
		digits = append(digits, int(d-'0'))
	}
	return digits, nil
}

// multipy returns the product of all values in nums.
func multipy(nums []int) int {
	p := 1
	for _, n := range nums {
		p *= n
	}
	return p
}

// max returns the max value from nums.
func max(nums []int) (m int) {
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return
}
