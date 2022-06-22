package embed_test

import "basics/embed"

func ExampleStruct() {
	embed.Struct()
	// Output:
	// GopherV2: embed.GopherV2{Gopher:embed.Gopher{Name:"Gala", Age:24, IsCoding:false, privateField:"ding doesn't change access modifiers"}, Friends:[]string{"Gabby", "Gael", "Garth", "Gazsi"}, Ratings:map[string]int{"chocolate":9, "coffee":3, "tea":7}}
	//
	// changed GopherV2: embed.GopherV2{Gopher:embed.Gopher{Name:"gopher", Age:26, IsCoding:true, privateField:"Can be changed because in same package."}, Friends:[]string{"Gabby", "Gael", "Garth", "Gazsi", "Gandalf"}, Ratings:map[string]int{"chocolate":9, "coffee":3, "garlic bread":10, "tea":7}}
	// Three layers of embedding {GopherV2:{Gopher:{Name:Ground Age:57005 IsCoding:true privateField:Access granted} Friends:[Gunter] Ratings:map[embedding:10]} Badge:129327}
}

func ExampleDeep() {
	embed.Deep()
	// Output:
	// {Residents:[{GopherV2:{Gopher:{Name:Gance Age:10 IsCoding:true privateField:Not accessible outside package.} Friends:[Guzz] Ratings:map[space:10]} Badge:128640} {GopherV2:{Gopher:{Name:Guuba Age:511 IsCoding:false privateField:Can be accessed in package.} Friends:[Ghorm Gokil] Ratings:map[death metal:10]} Badge:127755} {GopherV2:{Gopher:{Name:Gerry Age:88 IsCoding:true privateField:Put stuff here} Friends:[Gaqlyn Gicard Gosemary] Ratings:map[naps:10]} Badge:127881} {GopherV2:{Gopher:{Name:Gustion Age:21 IsCoding:true privateField:that you need to use} Friends:[Gidea Gno] Ratings:map[carbs:10]} Badge:127842} {GopherV2:{Gopher:{Name:Guna Age:255 IsCoding:false privateField:but other packages don't.} Friends:[Gouda] Ratings:map[haircuts:3]} Badge:127772}]}
	//
	// Not accessible outside package. Can be accessed in package. Put stuff here that you need to use but other packages don't.
}

func ExampleInterface() {
	embed.Interface(embed.Gopher{})
	// Output:
	// It was a dark and rainy night. The moon 🌕 felt so bright.
	// A strange passerby. Avert my gaze, I try.
	// Hi-dilly-ho, neighborinos!
	// The person yelled to me; suddenly something broke free.
	// Rawr XD
	// I ran as fast I could, through the streets to the forests of wood.
	// A final leap over some mud. The thing ceased with a great thud.
	// Rolled around for 20 mins
}

func ExampleInterface_v2() {
	embed.Interface(embed.GopherV2{})
	// Output:
	// It was a dark and rainy night. The moon 🌕 felt so bright.
	// A strange passerby. Avert my gaze, I try.
	// Hi-dilly-ho, neighborinos!
	// The person yelled to me; suddenly something broke free.
	// GOPHERV2 GRRRWAUGH!!
	// I ran as fast I could, through the streets to the forests of wood.
	// A final leap over some mud. The thing ceased with a great thud.
	// V2:Sat in mud for 20 mins
}

func ExampleInterface_v3() {
	embed.Interface(embed.GopherV3{})
	// Output:
	// It was a dark and rainy night. The moon 🌕 felt so bright.
	// A strange passerby. Avert my gaze, I try.
	// GOPHERV3 SAYS WHADDUP!!
	// The person yelled to me; suddenly something broke free.
	// GOPHERV2 GRRRWAUGH!!
	// I ran as fast I could, through the streets to the forests of wood.
	// A final leap over some mud. The thing ceased with a great thud.
	// V2:Sat in mud for 20 mins
}
