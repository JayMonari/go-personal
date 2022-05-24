package pnr_test

import (
	pnr "basics/panic-recover"
)

func ExamplePanicAfterDefer() {
	// XXX: No way to test output of a panic. That's **not** something you can
	// make an example out of, but we can at least have one to look at by
	// uncommenting 👇 the line below.
	// pnr.PanicAfterDefer()
	// XXX: Would be Output:
	// defer: Still print statement even with a panic
	// panic: 💣 TIME TO BLOW UP!!!
}

func ExamplePanicKeepCalm() {
	pnr.PanicKeepCalm()
	// Output:
	// recovered from: 😱 AWWW 💩WE'RE GOING DOWN!
}

func ExamplePanicNilPointer() {
	// XXX: No way to test output of a panic. That's **not** something you can
	// make an example out of, but we can at least have one to look at by
	// uncommenting 👇 the line below.
	// pnr.PanicAnatomy()
	// XXX: Would be Output:
	// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	//         panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4ef6f6]
	//
	// goroutine 1 [running]:
	// testing.(*InternalExample).processRunResult(0xc000057c68, {0x0, 0x0}, 0xc00007eb60?, 0x0, {0x5027e0, 0x5ef9f0})
	//         /usr/lib/go/src/testing/example.go:91 +0x4e5
	// testing.runExample.func2()
	//         /usr/lib/go/src/testing/run_example.go:59 +0x11c
	// panic({0x5027e0, 0x5ef9f0})
	//         /usr/lib/go/src/runtime/panic.go:838 +0x207
	// basics/panic-recover.(*myStruct).CausePanic(...)
	//         /home/jay/basics/panic-recover/panic_recover.go:25
	// basics/panic-recover.PanicAnatomy()
	//         /home/jay/basics/panic-recover/panic_recover.go:30 +0x16
	// basics/panic-recover_test.ExamplePanicAnatomy()
	//         /home/jay/basics/panic-recover/example_test.go:18 +0x17
	// testing.runExample({{0x5201b1, 0x13}, 0x527dd0, {0x0, 0x0}, 0x0})
	//         /usr/lib/go/src/testing/run_example.go:63 +0x28d
	// testing.runExamples(0xc000057e58, {0x5f4360?, 0x3, 0x0?})
	//         /usr/lib/go/src/testing/example.go:44 +0x186
	// testing.(*M).Run(0xc00010c140)
	//         /usr/lib/go/src/testing/testing.go:1721 +0x689
	// main.main()
	//         _testmain.go:53 +0x1aa
	// exit status 2
}
