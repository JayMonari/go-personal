package interfaces_test

import (
	"basics/interfaces"
	"fmt"
)

func ExampleFlyer() {
	_ = interfaces.Walker(interfaces.Duck{})
}

func ExampleGoForWalk() {
	m := interfaces.Man{}
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

func ExampleVisitWaterPark() {
	interfaces.VisitWaterPark(interfaces.Man{}, interfaces.Duck{})
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

type Seal struct{}

func (s Seal) Swim() { fmt.Println("Swims up to the surface. Arr! Arr! 🦭") }

func ExampleInterfacesToConcreteType() {
	fmt.Println("Put in Duck")
	interfaces.InterfacesToConcreteType(interfaces.Duck{IsFlying: true})
	fmt.Println("Put in Man")
	interfaces.InterfacesToConcreteType(interfaces.Man{})
	fmt.Println("Put in our own Swimmer (Seal)")
	interfaces.InterfacesToConcreteType(Seal{})
	// Output:
	// Put in Duck
	// Looks like this is a Duck! 🦆
	// And it is flying.
	// The duck 🦆 flies up.
	// The duck 🦆 waddles forward.
	// The duck 🦆 paddles around.
	// Put in Man
	// This isn't any type of Duck I've ever seen....
	// Time for takeoff 🛫
	// I'm walking, 🚶 yes indeed!
	// Splish Splash 🌊
	// Put in our own Swimmer (Seal)
	// This isn't any type of Duck I've ever seen....
	// This type interfaces_test.Seal doesn't have a mapping in the switch
}
