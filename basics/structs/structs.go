package structs

import "fmt"

type gopher struct {
	name     string
	age      int
	isCoding bool
}

type city struct {
	gophers         []gopher
	gopherAddresses map[gopher]string
}

// New is a constructor of a gopher, since New is exported (because it is
// capitalized) we can call it outside of the package, while keeping everything
// about a gopher internal!
func New(name string, age int, isCoding bool) gopher {
	return gopher{
		name:     name,
		age:      age,
		isCoding: isCoding,
	}
}

// StructBasic shows you how to initialize structs, manipulate all the values
// within a struct and use them in other structs.
func StructBasic() {
	var zero gopher
	gordo := gopher{name: "Gordo", age: 22, isCoding: true}
	gary := gopher{name: "Gary"}
	anon := gopher{age: 42, isCoding: true}
	fmt.Printf("zero valued gopher: %#v\n", zero)
	fmt.Printf("gordo gopher: %#v\n", gordo)
	fmt.Printf("gary gopher: %#v\n", gary)
	fmt.Printf("anon gopher: %#v\n", anon)
	fmt.Println()

	teska := city{
		gophers:         []gopher{gordo, gary, anon},
		gopherAddresses: map[gopher]string{gordo: "123 Lemon Dr.", gary: "889 Galaway Ave.", anon: "543 W 8th St."},
	}
	fmt.Printf("gopher city: %#v", teska)
	fmt.Println()

	gary.age = 33
	fmt.Printf("gary gopher: %#v\n", gary)
	anon.name = "Garfunkel"
	fmt.Printf("anon gopher: %#v\n", anon)

	gordo.age = 0
	gordo.name = ""
	gordo.isCoding = false
	fmt.Printf("gordo gopher: %#v\n", gordo)
}
