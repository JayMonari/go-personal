package series

// All returns a list of all substrings of s with length n.
func All(length int, digits string) []string {
	series := []string{}
	for i, j := 0, length; j < len(digits) + 1; i, j = i+1, j+1 {
		series = append(series, digits[i:j])
	}
	return series
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(length int, digits string) string {
	return digits[:length]
}

// First returns the first substring of digits with length if length is less
// than or equal to the length of the given string else it returns an empty
// string and false.
func First(length int, digits string) (first string, ok bool) {
	if length > len(digits) {
		return "", false
	}
	return digits[:length], true
}
