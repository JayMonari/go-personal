package allyourbase

import "errors"

// ConvertToBase returns a new slice of converted digits from the input base newBase
// the output base with the input slice. If all digits are not less than the
// inputBase an error is given back or if either base is less than 2.
func ConvertToBase(inBase int, digits []int, outBase int) ([]int, error) {
	switch {
	case inBase < 2:
		return []int{}, errors.New("input base must be >= 2")
	case outBase < 2:
		return []int{}, errors.New("output base must be >= 2")
	}

	baseValue := 0
	for _, d := range digits {
		if d < 0 || d >= inBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		baseValue = baseValue * inBase + d
	}

	return convert(baseValue, outBase), nil
}

// convert returns a slice of ints with each digit of the value converted to
// the newBase. If the value is 0 []int{0} is returned.
func convert(value, newBase int) []int {
	converted := []int{}
	for n := value; n > 0; n /= newBase {
		converted = append(converted, n % newBase)
	}
	if len(converted) == 0 {
		converted = append(converted, 0)
	}
	reverse(&converted)
	return converted
}

// reverse returns the passed in slice with the order of the elements reversed.
func reverse(sl *[]int) {
	s := *sl
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
