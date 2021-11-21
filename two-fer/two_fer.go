// Package twofer only implements one method ShareWith which shares one
// thing with another person
package twofer

import "fmt"

// ShareWith returns the string "One for <name>, one for me." If the name is
// empty it will return "you" as the name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
