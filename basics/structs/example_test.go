package structs_test

import (
	"basics/structs"
	"fmt"
)

func ExampleNew() {
	fmt.Printf("%#v", structs.New("Jay", 29, true))
	// Output: structs.gopher{name:"Jay", age:29, isCoding:true}
}

func ExampleStructBasic() {
	structs.StructBasic()
	// Output:
	// zero valued gopher: structs.gopher{name:"", age:0, isCoding:false}
	// gordo gopher: structs.gopher{name:"Gordo", age:22, isCoding:true}
	// gary gopher: structs.gopher{name:"Gary", age:0, isCoding:false}
	// anon gopher: structs.gopher{name:"", age:42, isCoding:true}
	//
	// gopher city: {gophers:[{name:Gordo age:22 isCoding:true} {name:Gary age:0 isCoding:false} {name: age:42 isCoding:true}] gopherAddresses:map[{name: age:42 isCoding:true}:543 W 8th St. {name:Gary age:0 isCoding:false}:889 Galaway Ave. {name:Gordo age:22 isCoding:true}:123 Lemon Dr.]}
	// gary gopher: structs.gopher{name:"Gary", age:33, isCoding:false}
	// anon gopher: structs.gopher{name:"Garfunkel", age:42, isCoding:true}
	// gordo gopher: structs.gopher{name:"", age:0, isCoding:false}
}
