package interfaces_test

import (
	"basics/interfaces"
	"fmt"
)

type Seal struct{}

func (s Seal) Swim() { fmt.Println("Swims up to the surface. Arr! Arr! 🦭") }

func ExampleWalker() {
	t := interfaces.Walker(interfaces.Duck{})
	// NOTE(jay): cannot convert (Seal literal) (value of type Seal) to
	// interfaces.Walker (Seal does not implement interfaces.Walker
	// (missing method Walk))
	// _ = interfaces.Walker(Seal{})
	fmt.Printf("%#v", t)
	// Output: interfaces.Duck{IsFlying:false}
}

func ExampleWalkSwimmer() {
	t := interfaces.WalkSwimmer(interfaces.Person(0))
	// NOTE(jay): cannot convert (Seal literal) (value of type Seal) to
	// interfaces.WalkSwimmer (Seal does not implement interfaces.WalkSwimmer
	// (missing method Walk))
	// _ = interfaces.WalkSwimmer(Seal{})
	fmt.Printf("%#v", t)
	// Output: 0x0
}

func ExampleGoForWalk() {
	m := interfaces.Person(0)
	d := interfaces.Duck{}
	interfaces.GoForWalk(m, d)
	// Output:
	// It was looking like a great day outside. ☀️
	// Two very different types decided to go for a walk.
	// I'm walking, 🚶 yes indeed!
	// The duck 🦆 waddles forward.
	// They ran into each other and locked eyes 👀. What will happen now?
	// I'm walking, 🚶 yes indeed!
	// The duck 🦆 waddles forward.
	// Looks like they decided to continue their walk together! 😄
}

func ExampleSoarIntoTheClouds() {
	interfaces.SoarIntoTheClouds(interfaces.Duck{})
	// Output:
	// The clouds ☁️ look so good today!
	// The duck 🦆 flies up.
	// Feels good to be on Cloud Nine. 😎
}

func ExampleInterfacesToConcreteType() {
	fmt.Println("Put in Duck")
	interfaces.InterfacesToConcreteType(interfaces.Duck{IsFlying: true})
	fmt.Println("Put in Person")
	interfaces.InterfacesToConcreteType(interfaces.Person(0))
	fmt.Println("Put in our own Swimmer (Seal)")
	interfaces.InterfacesToConcreteType(Seal{})
	// Output:
	// Put in Duck
	// Looks like this is a Duck! 🦆
	// And it is flying.
	// The duck 🦆 flies up.
	// The duck 🦆 waddles forward.
	// The duck 🦆 paddles around.
	// Put in Person
	// This isn't any type of Duck I've ever seen....
	// Time for takeoff 🛫
	// I'm walking, 🚶 yes indeed!
	// Splish Splash 🌊
	// Put in our own Swimmer (Seal)
	// This isn't any type of Duck I've ever seen....
	// This type interfaces_test.Seal doesn't have a mapping in the switch
}

func ExampleVisitWaterPark() {
	interfaces.VisitWaterPark(interfaces.Person(0), interfaces.Duck{})
	// Output:
	// Two very different types decided to go to a water park.
	// I'm walking, 🚶 yes indeed!
	// The duck 🦆 waddles forward.
	// They both make it in and find a pool to dive into.
	// The duck 🦆 paddles around.
	// Splish Splash 🌊
	// Uh oh, looks like interfaces.Duck didn't like that!
	// The duck 🦆 flies up.
}
