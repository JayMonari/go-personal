package etl

import "strings"

// Transform takes all values from each key and flattens them out as keys to a
// new map with the value of the original key.
// e.g. {1: {"A", "B"}} -> {"A": 1, "B": 1}
func Transform(data map[int][]string) map[string]int {
	t := make(map[string]int, len(data))
	for k, v := range data {
		for _, s := range v {
			t[strings.ToLower(s)] = k
		}
	}
	return t
}
