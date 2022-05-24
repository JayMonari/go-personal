package pnr

import (
	"fmt"
)

func PanicAfterDefer() {
	defer fmt.Println("defer: Still print statement even with a panic")
	panic("💣 TIME TO BLOW UP!!!")
}

func PanicKeepCalm() {
	defer recuperate()
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
	s = nil
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
	// panic: runtime error: index out of range [4] with length 4
	//
	// goroutine 1 [running]:
	// main.main()
	//         /home/jay/basics/cmd/main.go:7 +0x1b
	// exit status 2
}
