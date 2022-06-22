package structs_test

import (
	"basics/structs"
	"fmt"
)

func ExampleNew() {
	literalGopher := structs.Gopher{
		Name:     "Gitral",
		Age:      0o703,
		IsCoding: true,
	}
	fmt.Printf("Can never set privateField: %#v\n", literalGopher)

	constructedGopher :=
		structs.New("Jay", 29, true, "once set, can't be changed.")
	fmt.Printf("%#v\n", constructedGopher)

	constructedGopher.Age = 58
	constructedGopher.IsCoding = false
	constructedGopher.Name = "Jöt"
	// NOTE(jay): Can't do!
	// constructedGopher.privateField = "Not possible!"
	fmt.Printf("%#v\n", constructedGopher)

	// Output:
	// Can never set privateField: structs.Gopher{Name:"Gitral", Age:451, IsCoding:true, privateField:""}
	// structs.Gopher{Name:"Jay", Age:29, IsCoding:true, privateField:"once set, can't be changed."}
	// structs.Gopher{Name:"Jöt", Age:58, IsCoding:false, privateField:"once set, can't be changed."}
}

func ExampleBasic() {
	city := structs.Basic()
	for _, g := range city.Gophers {
		if g.Name == "Garfunkel" {
			fmt.Printf("We found him: %s\n", city.GopherAddresses[g])
		}
	}
	// Output:
	// zero valued gopher: structs.Gopher{Name:"", Age:0, IsCoding:false, privateField:""}
	// gordo gopher: structs.Gopher{Name:"Gordo", Age:22, IsCoding:true, privateField:"Set it and forget it"}
	// gary gopher: structs.Gopher{Name:"Gary", Age:0, IsCoding:false, privateField:""}
	// anon gopher: structs.Gopher{Name:"", Age:42, IsCoding:true, privateField:"Scanning 60000 ports"}
	//
	// gary gopher: structs.Gopher{Name:"Gary", Age:33, IsCoding:false, privateField:"Searching: Why does my husband fart so much."}
	// anon gopher: structs.Gopher{Name:"Garfunkel", Age:42, IsCoding:true, privateField:"Scanning 60000 ports"}
	// Gordo is in the middle of coding! Come back soon.
	// {Gary 33 false Searching: Why does my husband fart so much.} lives at 889 Galaway Ave.
	// Garfunkel is in the middle of coding! Come back soon.
	//
	// gordo gopher: structs.Gopher{Name:"", Age:0, IsCoding:false, privateField:""}
	// zero  gopher: structs.Gopher{Name:"", Age:0, IsCoding:false, privateField:""}
	// We found him: 543 W 8th St.
}
