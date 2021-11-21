package accumulate

// Accumulate takes in a slice of strings and returns a new slice transformed
// by the converter function provided.
func Accumulate(words []string, converter func(string) string) []string {
	transformed := []string{}
	for _, word := range words {
		transformed = append(transformed, converter(word))
	}
	return transformed
}
