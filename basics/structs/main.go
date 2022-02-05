package main

import "fmt"

func main() {
	StructBasic()
	StructEmbed()
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

func StructBasic() {
	gordo := gopher{name: "Gordo", age: 22, isCoding: true}
	gary := gopher{name: "Gary"}
	anon := gopher{age: 42, isCoding: true}
	fmt.Printf("gordo gopher: %#v\n", gordo)
	fmt.Printf("gary gopher: %#v\n", gary)
	fmt.Printf("anon gopher: %#v\n", anon)
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

func StructEmbed() {
	friends := []string{"Gabby", "Gael", "Garth", "Gazsi"}
	myRatings := map[string]int{"coffee": 3, "tea": 7, "chocolate": 9}
	gala := gopher2{
		gopher:    gopher{name: "Gala", age: 24, isCoding: false},
		friends:   friends,
		myRatings: myRatings,
	}
	fmt.Printf("gopher2: %#v\n", gala)
	fmt.Println()

	gala.name = "gopher"
	gala.age = 26
	gala.isCoding = true
	gala.friends = append(gala.friends, "Gandalf")
	gala.myRatings["garlic bread"] = 10
	fmt.Printf("changed gopher2: %#v\n", gala)
}
