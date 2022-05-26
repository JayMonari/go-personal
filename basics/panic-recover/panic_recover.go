package pnr

import (
	"fmt"
)

// PanicAfterDefer shows that even if a panic occurs in a function a `defer`
// statement will **always** execute. This is to make sure system resources are
// cleaned up and why we can `recover` in the first place.
func PanicAfterDefer() {
	defer fmt.Println("defer: Still print statement even with a panic")
	panic("💣 TIME TO BLOW UP!!!")
}

// PanicKeepCalm shows how we can `recover` from a `panic` by using a `defer`
// statement that calls `recover()`. You **must** put recover in a `defer`
// statement or else it won't work.
func PanicKeepCalm() {
	defer recuperate()
	// XXX: This will not stop the panic
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

func PanicNilPointer() {
	s := new(myStruct)
	s = nil // XXX: Obviously dangerous, but it happens in mysterious ways.
	fmt.Println(s.CausePanic())
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x47df56]
	//
	// goroutine 1 [running]:
	// main.(*myStruct).CausePanic(...)
	// 	/home/jay/basics/cmd/main.go:7
	// main.main()
	// 	/home/jay/basics/cmd/main.go:12 +0x16
	// exit status 2
}

func PanicIndexOut() {
	daBomb := []string{"set", "us", "up", "da bomb."}
	fmt.Println(daBomb[len(daBomb)])
	// We actually want:
	// fmt.Println(daBomb[len(daBomb)-1])
}
