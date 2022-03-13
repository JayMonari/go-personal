package interfaces

import "fmt"

type Walker interface{ Walk() }

type Swimmer interface{ Swim() }

type Flyer interface{ Fly() }

type WalkSwimmer interface {
	Walker
	Swimmer
}

type WalkSwimFlyer interface {
	Walker
	Swimmer
	Flyer
}

type Man struct{}

func (m Man) Walk() { fmt.Println("I'm walking, 🚶 yes indeed!") }
func (m Man) Swim() { fmt.Println("Splish Splash 🌊") }
func (m Man) Fly()  { fmt.Println("Time for takeoff 🛫") }

type Duck struct{ IsFlying bool }

func (d Duck) Walk() { fmt.Println("The duck 🦆 waddles forward.") }
func (d Duck) Swim() { fmt.Println("The duck 🦆 paddles around.") }
func (d Duck) Fly()  { fmt.Println("The duck 🦆 flies up.") }

// GoForWalk is an example of using an interface to satisfy a condition we want
// our parameters (man and duck) to have, without forcing someone to have
// an exact implementations
func GoForWalk(man Walker, duck Walker) {
	fmt.Println("It was looking like a great day outside. ☀️")
	fmt.Println("Two very different types decided to go for a walk.")
	man.Walk()
	duck.Walk()
	fmt.Println("They ran into each other and locked eyes 👀. What will happen now?")
	man.Walk()
	duck.Walk()
	fmt.Println("Looks like they decided to continue their walk together! 😄")
}

// SoarIntoTheClouds will take the Flyer and put them sky high into the clouds.
func SoarIntoTheClouds(duck Flyer) {
	fmt.Println("The clouds ☁️ look so good today!")
	duck.Fly()
	// XXX: Notice we **cannot** call the ducks other methods
	//  (type Flyer has no field or method Walk/Swim)
	// duck.Walk()
	// duck.Swim()
	fmt.Println("Feels good to be on Cloud Nine. 😎")
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

func InterfacesToConcreteType(s Swimmer) {
	// XXX: s.isFlying undefined (type Swimmer has no field or method isFlying)
	// This shows us that even if Duck had more methods or had any fields we only
	// can use what is satisfied by the interface.
	// s.isFlying
	d, ok := s.(Duck)
	if ok {
		fmt.Println("Looks like this is a Duck! 🦆")
		if d.IsFlying {
			fmt.Println("And it is flying.")
		} else {
			fmt.Println("And it isn't flying.")
		}
	} else {
		fmt.Println("This isn't any type of Duck I've ever seen....")
	}
	// If we are unsure of the type we can use a switch type statement from a
	// previous lesson!
	switch t := s.(type) {
	case WalkSwimFlyer:
		t.Fly()
		t.Walk()
		t.Swim()
	case Duck:
		if !d.IsFlying {
			d.IsFlying = true
		}
		d.Fly()
	default:
		fmt.Printf("This type %T doesn't have a mapping in the switch\n", t)
	}
}
