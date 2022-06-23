package structs

import "fmt"

// Gopher is a public struct that can be made outside of the package. It
// consists of some basic fields that all gophers have and a privateField that
// can only be accessed in the package. We make things private or better said
// -- unexported -- so the people using our library (this code) don't have to
// worry about certain fields.
type Gopher struct {
	Name         string
	Age          int
	IsCoding     bool
	privateField string
}

// city is a unexported struct with a struct inside of it! 🤯 That's because a
// gopher is a `type` and we can put ALL types into a struct. That means we can
// put city into another struct called state and it would haves cities in it
// with gophers in them! 🤯
type city struct {
	Gophers         []Gopher
	GopherAddresses map[Gopher]string
}

// New is a constructor of a Gopher, since New is exported (because it is
// capitalized) we can call it outside of the package, while being able to set
// private fields of our Gopher!
func New(name string, age int, isCoding bool, privateField string) Gopher {
	return Gopher{
		Name:         name,
		Age:          age,
		IsCoding:     isCoding,
		privateField: privateField,
	}
}

// Basic shows you how to initialize (make) structs, manipulate all the
// values within a struct by getting and setting the values and use them in
// other structs. We also return a `city` struct here to show you can give back
// unexported types from exported functions.
func Basic() city {
	// Make a gopher and have ALL fields set to the zero value.
	var zero Gopher
	// Make a gopher and set all fields to what we want them to be.
	gordo := Gopher{
		Name:         "Gordo",
		Age:          22,
		IsCoding:     true,
		privateField: "Set it and forget it",
	}
	// Make a gopher and only set the fields we care about, leaving the rest to
	// be initialized (made) with their zero values.
	gary := Gopher{Name: "Gary"}
	anon := Gopher{Age: 42, IsCoding: true, privateField: "Scanning 60000 ports"}
	fmt.Printf("zero valued gopher: %#v\n", zero)
	fmt.Printf("gordo gopher: %#v\n", gordo)
	fmt.Printf("gary gopher: %#v\n", gary)
	fmt.Printf("anon gopher: %#v\n", anon)
	fmt.Println()

	// Access a value by using the `.` and the fields name
	gary.Age = 33
	gary.privateField = "Searching: Why does my husband fart so much."
	fmt.Printf("gary gopher: %#v\n", gary)
	anon.Name = "Garfunkel"
	fmt.Printf("anon gopher: %#v\n", anon)

	teska := city{
		Gophers: []Gopher{gordo, gary, anon},
		GopherAddresses: map[Gopher]string{
			gordo: "123 Lemon Dr.",
			gary:  "889 Galaway Ave.",
			anon:  "543 W 8th St.",
		},
	}
	// Since teska has a slice of gophers we can get it and range over each of
	// them in a for loop. g == gopher
	for _, g := range teska.Gophers {
		// Access each gopher's IsCoding field. In the slice of gophers we are
		// accessing from the city!
		if g.IsCoding {
			fmt.Println(g.Name, "is in the middle of coding! Come back soon.")
			continue
		}
		fmt.Println(g, "lives at", teska.GopherAddresses[g])
	}
	fmt.Println()

	// zero out a gopher, not needed here, but you can see how it is done.
	gordo.Age = 0
	gordo.Name = ""
	gordo.IsCoding = false
	gordo.privateField = ""
	fmt.Printf("gordo gopher: %#v\nzero  gopher: %#v\n", gordo, zero)
	return teska
}
