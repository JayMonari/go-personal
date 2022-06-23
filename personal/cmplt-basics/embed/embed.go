package embed

import "fmt"

// Gopher is an example struct to show off embedding
type Gopher struct {
	Name     string
	Age      int
	IsCoding bool

	privateField string
}

// GopherV2 has a Gopher embedded inside of it and has some additional fields.
type GopherV2 struct {
	// We can make this private if we want like:
	// `gopher Gopher`
	// but that would remove the inner type promotion, meaning we can't do fun
	// and easy stuff like `myGopherV2.Name` 👍, we would have to do
	// `myGopherV2.gopher.Name` 👎
	Gopher

	Friends []string
	Ratings map[string]int
}

// GopherV3 has a GopherV2 which has a Gopher in it, so it is allowed to access
// all the fields and methods of both GopherV2 and Gopher along with it's own
// fields and methods.
type GopherV3 struct {
	GopherV2
	Badge rune
}

// Struct is an example function that shows how to initialize both
// versions of embedded structs and how to access the embedded fields.
func Struct() {
	gala := GopherV2{
		Gopher: Gopher{
			Name:         "Gala",
			Age:          24,
			IsCoding:     false,
			privateField: "ding doesn't change access modifiers"},
		Friends: []string{"Gabby", "Gael", "Garth", "Gazsi"},
		Ratings: map[string]int{"coffee": 3, "tea": 7, "chocolate": 9},
	}
	fmt.Printf("GopherV2: %#v\n", gala)
	fmt.Println()

	// Change our GopherV2
	gala.Name = "gopher"
	gala.Age = 26
	gala.IsCoding = true
	gala.privateField = "Can be changed because in same package."
	gala.Friends = append(gala.Friends, "Gandalf")
	gala.Ratings["garlic bread"] = 10
	fmt.Printf("changed GopherV2: %#v\n", gala)

	// This is here to show that you **always** need to have `field: field{}`
	// when making your struct inline, but don't need to when accessing fields
	// e.g. `g.name` works! Instead of `g.GopherV2.Gopher.name`
	g := GopherV3{
		GopherV2: GopherV2{
			Gopher: Gopher{
				Name:         "Ground",
				Age:          0xdead,
				IsCoding:     true,
				privateField: "Access granted"},
			Friends: []string{"Gunter"},
			Ratings: map[string]int{"embedding": 10},
		},
		Badge: '🤯',
	}
	fmt.Printf("Three layers of embedding %+v\n", g)
}

// City is a struct showing structs with embedded fields doesn't change
// anything about the way we expect a struct to behave
type City struct{ Residents []GopherV3 }

// Deep shows that structs with embedded fields act no differently with
// the embedded fields.
func Deep() {
	c := City{
		Residents: []GopherV3{
			{ // 👈 Notice we don't need to put GopherV3 here
				GopherV2: GopherV2{
					Gopher: Gopher{
						Name:         "Gance",
						Age:          0b001010,
						IsCoding:     true,
						privateField: "Not accessible outside package."},
					Friends: []string{"Guzz"},
					Ratings: map[string]int{"space": 10},
				},
				Badge: '🚀',
			},
			{ // 👈 Notice we don't need to put GopherV3 here
				GopherV2: GopherV2{
					Gopher: Gopher{
						Name:         "Guuba",
						Age:          0o777,
						IsCoding:     false,
						privateField: "Can be accessed in package."},
					Friends: []string{"Ghorm", "Gokil"},
					Ratings: map[string]int{"death metal": 10},
				},
				Badge: '🌋',
			},
			{ // 👈 Notice we don't need to put GopherV3 here
				GopherV2: GopherV2{
					Gopher: Gopher{
						Name:         "Gerry",
						Age:          88,
						IsCoding:     true,
						privateField: "Put stuff here"},
					Friends: []string{"Gaqlyn", "Gicard", "Gosemary"},
					Ratings: map[string]int{"naps": 10},
				},
				Badge: '🎉',
			},
			{ // 👈 Notice we don't need to put GopherV3 here
				GopherV2: GopherV2{
					Gopher: Gopher{
						Name:         "Gustion",
						Age:          21,
						IsCoding:     true,
						privateField: "that you need to use"},
					Friends: []string{"Gidea", "Gno"},
					Ratings: map[string]int{"carbs": 10},
				},
				Badge: '🍢',
			},
			{ // 👈 Notice we don't need to put GopherV3 here
				GopherV2: GopherV2{
					Gopher: Gopher{
						Name:         "Guna",
						Age:          0b11111111,
						IsCoding:     false,
						privateField: "but other packages don't."},
					Friends: []string{"Gouda"},
					Ratings: map[string]int{"haircuts": 3},
				},
				Badge: '🌜',
			},
		},
	}
	fmt.Printf("%+v\n\n", c)
	for _, g := range c.Residents {
		// g.GopherV2.Gopher.privateField
		fmt.Printf("%s ", g.privateField)
	}
}

type Human interface{ Greet() string }

type Bear interface{ Growl() string }

type Pig interface{ MudBath(minutes int) }

// HumanBearPig is the thing of nightmares.
type HumanBearPig interface {
	Human // half human
	Bear  // half bear
	Pig   // half pig
}

func (g Gopher) Greet() string { return "Hi-dilly-ho, neighborinos!" }
func (g Gopher) Growl() string { return "Rawr XD" }
func (g Gopher) MudBath(m int) { fmt.Printf("Rolled around for %d mins", m) }

func (g GopherV2) Growl() string { return "GOPHERV2 GRRRWAUGH!!" }
func (g GopherV2) MudBath(m int) { fmt.Printf("V2:Sat in mud for %d mins", m) }

func (g GopherV3) Greet() string { return "GOPHERV3 SAYS WHADDUP!!" }

// Interface shows that we need a struct that will satisfy the entire
// interface, but we don't care if that struct or the embedded structs inside
// of it satisfy the interface.
func Interface(hbp HumanBearPig) {
	fmt.Println("It was a dark and rainy night. The moon 🌕 felt so bright.")
	fmt.Println("A strange passerby. Avert my gaze, I try.")
	fmt.Println(hbp.Greet())
	fmt.Println("The person yelled to me; suddenly something broke free.")
	fmt.Println(hbp.Growl())
	fmt.Println("I ran as fast I could, through the streets to the forests of wood.")
	fmt.Println("A final leap over some mud. The thing ceased with a great thud.")
	hbp.MudBath(20)
}
