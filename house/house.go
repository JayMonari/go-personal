package house

import (
	"fmt"
	"strings"
)

// items are all the nouns of the nursery rhyme.
var items = [12]string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

// actions are the verbs for each item.
var actions = [12]string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
	"",
}

// Verse produces the associated verse depending on the verse number n.
func Verse(n int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("This is %s", items[n-1]))
	for n = n - 2; n >= 0; n-- {
		sb.WriteString(fmt.Sprintf("\nthat %s %s", actions[n], items[n]))
	}
	return sb.String()
}

// Song returns the entire nursery rhyme.
func Song() string {
	vs := make([]string, 12)
	for n := 1; n <= 12; n++ {
		vs[n-1] = Verse(n)
	}
	return strings.Join(vs, "\n\n")
}
