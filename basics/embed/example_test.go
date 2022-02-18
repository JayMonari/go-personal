package embed_test

import "basics/embed"

func ExampleStructEmbed() {
	embed.StructEmbed()
	// Output:
	// gopher2: embed.gopher2{gopher:embed.gopher{name:"Gala", age:24, isCoding:false}, friends:[]string{"Gabby", "Gael", "Garth", "Gazsi"}, myRatings:map[string]int{"chocolate":9, "coffee":3, "tea":7}}
	//
	// changed gopher2: embed.gopher2{gopher:embed.gopher{name:"gopher", age:26, isCoding:true}, friends:[]string{"Gabby", "Gael", "Garth", "Gazsi", "Gandalf"}, myRatings:map[string]int{"chocolate":9, "coffee":3, "garlic bread":10, "tea":7}}
}
