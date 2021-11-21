package allergies

// items is a map of all allergens and their given score.
var items = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns a slice of strings of all items that a patient is allergic
// to based on the given score.
func Allergies(score uint) []string {
	a := []string{}
	for item, val := range items {
		if score&val != 0 {
			a = append(a, item)
		}
	}
	return a
}

// AllergicTo returns whether or not a given item is within the score.
func AllergicTo(score uint, item string) bool {
	if score&items[item] != 0 {
		return true
	}
	return false
}
