package pnr

import (
	"fmt"
)

// AfterDefer shows that even if a panic occurs in a function a `defer`
// statement will **always** execute. This is to make sure system resources are
// cleaned up and why we can `recover` in the first place.
func AfterDefer() {
	defer fmt.Println("defer: Still print statement even with a panic")
	panic("💣 TIME TO BLOW UP!!!")
}

// KeepCalm shows how we can `recover` from a `panic` by using a `defer`
// statement that calls `recover()`. You **must** put recover in a `defer`
// statement or else it won't work.
func KeepCalm() {
	defer recuperate()
	// NOTE(jay): This will not stop the panic
	// recover()
	panic("😱 AWWW 💩WE'RE GOING DOWN!")
}

func recuperate() {
	if err := recover(); err != nil {
		fmt.Println("recovered from:", err)
	}
}

type myStruct struct{ cantAccess string }

func (s *myStruct) CausePanic() string { return s.cantAccess }

func NilPointer() {
	s := new(myStruct)
	s = nil // NOTE(jay): Obviously dangerous, but it happens in mysterious ways.
	fmt.Println(s.CausePanic())
}

func NewMap() {
	m := new(map[string]string)
	(*m)["nil map"] = "causes panic!"
	// We actually want:
	// ma := make(map[string]string)
	// ma["not nil"] = "works"
}

func IndexOut() {
	daBomb := []string{"set", "us", "up", "da bomb."}
	fmt.Println(daBomb[len(daBomb)])
	// We actually want:
	// fmt.Println(daBomb[len(daBomb)-1])
}
