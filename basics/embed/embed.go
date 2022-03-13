package embed

import "fmt"

// TODO(jaymonari): Interface func example and embedded struct makes container
// struct satisfy an interface.

type Man interface {
	Greet() string
}

type Bear interface {
	Growl() string
}

type Pig interface {
	Squeal() string
}

type ManBearPig interface {
	Man  // half man
	Bear // half bear
	Pig  // half pig
}

type gopher struct {
	name     string
	age      int
	isCoding bool
}

type gopher2 struct {
	gopher

	friends   []string
	myRatings map[string]int
}

type gopher3 struct {
	gopher2

	badge rune
}

func StructEmbed() {
	gala := gopher2{
		gopher:    gopher{name: "Gala", age: 24, isCoding: false},
		friends:   []string{"Gabby", "Gael", "Garth", "Gazsi"},
		myRatings: map[string]int{"coffee": 3, "tea": 7, "chocolate": 9},
	}
	fmt.Printf("gopher2: %#v\n", gala)
	fmt.Println()

	// Change our gopher2
	gala.name = "gopher"
	gala.age = 26
	gala.isCoding = true
	gala.friends = append(gala.friends, "Gandalf")
	gala.myRatings["garlic bread"] = 10
	fmt.Printf("changed gopher2: %#v\n", gala)

	// Hopefully you will never have to do this.
	// This is here to show that you **always** need to have `field: field{}`
	// when making your struct inline, but don't need to when accessing fields
	// e.g. `g.name` works! Instead of `g.gopher2.gopher.name`
	g := gopher3{
		gopher2: gopher2{
			gopher:    gopher{name: "Deep", age: 0xdead},
			friends:   []string{},
			myRatings: map[string]int{"embedding": 10},
		},
		badge: '🤯',
	}
	fmt.Printf("Three layers of embedding %+v\n", g)
}
