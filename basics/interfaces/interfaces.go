package interfaces

import (
	"fmt"
)

// Walker is the interface that wraps the basic Walk method.
type Walker interface{ Walk() }

// Swimmer is the interface that wraps the basic Swim method.
type Swimmer interface{ Swim() }

// Flyer is the interface that wraps the basic Fly method.
type Flyer interface{ Fly() }

// WalkSwimmer is the interface that groups the basic Walk and Swim methods.
type WalkSwimmer interface {
	Walker
	Swimmer
}

// WalkSwimmerFlyer is the interface that groups the basic Walk, Swim and Fly
// methods.
type WalkSwimFlyer interface {
	Walker
	Swimmer
	Flyer
}

// Person is a new type and it can satisfy all of our interfaces, the type system
// in Go is very powerful for features like this.
type Person uint8

func (p Person) Walk() { fmt.Println("I'm walking, 🚶 yes indeed!") }
func (p Person) Swim() { fmt.Println("Splish Splash 🌊") }
func (p Person) Fly()  { fmt.Println("Time for takeoff 🛫") }

// Duck is a struct type that satisfies all of our interfaces as well. We can
// see it has some extra information on whether or not it is flying.
type Duck struct{ IsFlying bool }

func (d Duck) Walk() { fmt.Println("The duck 🦆 waddles forward.") }
func (d Duck) Swim() { fmt.Println("The duck 🦆 paddles around.") }
func (d Duck) Fly()  { fmt.Println("The duck 🦆 flies up.") }

// GoForWalk is an example of using an interface to satisfy a condition we want
// our parameters (person and duck) to have, without forcing someone to have
// exact implementations. So they could put in anything that can walk.
func GoForWalk(person Walker, duck Walker) {
	fmt.Println("It was looking like a great day outside. ☀️")
	fmt.Println("Two very different types decided to go for a walk.")
	person.Walk()
	duck.Walk()
	fmt.Println("They ran into each other and locked eyes 👀. What will happen now?")
	person.Walk()
	duck.Walk()
	fmt.Println("Looks like they decided to continue their walk together! 😄")
}

// SoarIntoTheClouds will take the Flyer and put them sky high into the clouds.
func SoarIntoTheClouds(f Flyer) {
	fmt.Println("The clouds ☁️ look so good today!")
	f.Fly()
<<<<<<< Updated upstream
	// XXX: Notice we **cannot** call the other duck or person methods
=======
 	// XXX: Notice we **cannot** call the other duck or person methods
>>>>>>> Stashed changes
	//  (type Flyer has no field or method Walk/Swim)
	// f.Walk()
	// f.Swim()
	fmt.Println("Feels good to be on Cloud Nine. 😎")
}

// InterfacesToConcreteType shows us how we would turn an interface that only
// knows the Swim method into its concrete type, which allows us to gain access
// to that type's other methods and fields.
func InterfacesToConcreteType(s Swimmer) {
	// XXX: s.isFlying undefined (type Swimmer has no field or method isFlying)
	// This shows us that even if Duck had more methods or had any fields we only
	// can use what is satisfied by the interface.
	// s.isFlying
	d, ok := s.(Duck)
	switch ok {
	case true:
		fmt.Println("Looks like this is a Duck! 🦆")
		if d.IsFlying {
			fmt.Println("And it is flying.")
		} else {
			fmt.Println("And it isn't flying.")
		}
	case false:
		fmt.Println("This isn't any type of Duck I've ever seen....")
	}
	// If we are unsure of the type we can use a switch type statement from a
	// previous lesson. Very useful for JSON responses.
	switch t := s.(type) {
	case WalkSwimFlyer:
		t.Fly()
		t.Walk()
		t.Swim()
	case Duck:
		if !t.IsFlying {
			t.IsFlying = true
		}
		t.Fly()
	default:
		fmt.Printf("This type %T doesn't have a mapping in the switch\n", t)
	}
}

// VisitWaterPark takes in a type that can both walk and swim **and** a type
// that can walk, swim, and fly. It uses all of these methods inside of the
// function and **cannot** use anymore methods than the declared one for that
// type.
func VisitWaterPark(ws WalkSwimmer, wsf WalkSwimFlyer) {
	fmt.Println("Two very different types decided to go to a water park.")
	ws.Walk()
	wsf.Walk()
	fmt.Println("They both make it in and find a pool to dive into.")
	wsf.Swim()
	ws.Swim()
	fmt.Printf("Uh oh, looks like %T didn't like that!\n", wsf)
	wsf.Fly()
}
