package flatten

// Flatten takes in a slice of nested values and flattens it into one slice of
// integer. If the nested slices has anything other than int it will be ignored
// in the final result.
func Flatten(nested interface{}) []interface{} {
	flattened := []interface{}{}
	switch unknown := nested.(type) {
	case []interface{}:
		for _, ele := range unknown {
			flattened = append(flattened, Flatten(ele)...)
		}
	case int:
		flattened = append(flattened, unknown)
	}
	return flattened
}
