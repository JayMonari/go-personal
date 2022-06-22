package method

import "fmt"

type Gopher struct {
	name      string
	isCoding  bool
	favNumber int
}

// New returns a constructed Gopher of the given values.
func New(name string, isCoding bool, favNumber int) *Gopher {
	return &Gopher{name: name, isCoding: isCoding, favNumber: favNumber}
}

// DoesNotChangeFavNumber is a value receiver method that does not change the
// favNumber of the Gopher.
func (g Gopher) DoesNotChangeFavNumber(favNumber int) {
	// Since `g` is a copy and not a pointer `*`, the favNumber we update is
	// specific to this copy of `g`.
	// The original `g` still has its favNumber intact.
	g.favNumber = favNumber
}

// DoesChangeFavNumber is a pointer receiver method that will change the
// favNumber of the Gopher.
func (g *Gopher) DoesChangeFavNumber(favNumber int) {
	// We can drop the dereference to the pointer `(*g)` to just `g` if we wanted
	// because Go will automagically 🪄 dereference pointer values for us!
	(*g).favNumber = favNumber
	// This works too and is preferred:
	// g.favNumber = favNumber
}

// StartCoding is a pointer receiver method that starts the Gopher to coding.
func (g *Gopher) StartCoding() { g.isCoding = true }

// StopCoding is a pointer receiver method that stops the Gopher from coding.
func (g *Gopher) StopCoding() { g.isCoding = false }

// StartCoding is a package function that starts the Gopher to coding.
func StartCoding(g *Gopher) { g.isCoding = true }

// StopCoding is a package function that stops the Gopher from coding.
func StopCoding(g *Gopher) { g.isCoding = false }

// String satisfies the fmt.Stringer interface and will be used anytime we try
// to fmt.Print[f,ln](g) -- g == Gopher.
// We can see it is a value receiver method, meaning it takes a copy of a
// Gopher. This is because we **don't** want to update any values of our
// original Gopher and by using a value receiver method we make sure that
// cannot happen.
func (g Gopher) String() string {
	if g.isCoding {
		return "Let me get back to you after I'm done coding."
	} else {
		return fmt.Sprint(
			"Hi! I'm ", g.name, " and my favorite number is ", g.favNumber)
	}
}
