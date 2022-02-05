package main

func ExampleStructBasic() {
	StructBasic()
	// Output:
	// gordo gopher: main.gopher{name:"Gordo", age:22, isCoding:true}
	// gary gopher: main.gopher{name:"Gary", age:0, isCoding:false}
	// anon gopher: main.gopher{name:"", age:42, isCoding:true}
	//
	// gary gopher: main.gopher{name:"Gary", age:33, isCoding:false}
	// anon gopher: main.gopher{name:"Garfunkel", age:42, isCoding:true}
	// gordo gopher: main.gopher{name:"", age:0, isCoding:false}
}

func ExampleStructEmbed() {
	StructEmbed()
	// Output:
	// gopher2: main.gopher2{gopher:main.gopher{name:"Gala", age:24, isCoding:false}, friends:[]string{"Gabby", "Gael", "Garth", "Gazsi"}, myRatings:map[string]int{"chocolate":9, "coffee":3, "tea":7}}
	//
	// changed gopher2: main.gopher2{gopher:main.gopher{name:"gopher", age:26, isCoding:true}, friends:[]string{"Gabby", "Gael", "Garth", "Gazsi", "Gandalf"}, myRatings:map[string]int{"chocolate":9, "coffee":3, "garlic bread":10, "tea":7}}
}
