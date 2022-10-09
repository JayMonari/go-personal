package main

import (
	"errors"
	"fmt"
	"strings"
	"testing/iotest"
)

func ExampleDataErrReader() {
	// Output:
	//
}

func ExampleErrReader() {
	iotest.ErrReader(errors.New("read error"))
	// Output:
	//
}

func ExampleHalfReader() {
	hrdr := iotest.HalfReader(strings.NewReader(`This is going to be cut in half.`))
	srdr := strings.NewReader(`This is going to be cut in half.`)
	data := make([]byte, 16)
	n := 0
	n, _ = srdr.Read(data)
	fmt.Printf("strings.Reader: amount read: %d, data: %q\n", n, data)
	n, _ = hrdr.Read(data)
	fmt.Printf("iotest.HalfReader: amount read: %d, data: %q\n", n, data)
	n, _ = srdr.Read(data)
	fmt.Printf("strings.Reader: amount read: %d, data: %q\n", n, data)
	n, _ = hrdr.Read(data)
	fmt.Printf("iotest.HalfReader: amount read: %d, data: %q\n", n, data)
	// Output:
	// strings.Reader: "This is going to"
	// iotest.HalfReader: "This is going to"
	// strings.Reader: " be cut in half."
	// iotest.HalfReader: "going toin half."
}

func ExampleNewReadLogger() {
	// Output:
	//
}

func ExampleNewWriteLogger() {
	// Output:
	//
}

func ExampleOneByteReader() {
	// Output:
	//
}

func ExampleTestReader() {
	// Output:
	//
}

func ExampleTimeoutReader() {
	// Output:
	//
}

func ExampleTruncateWriter() {
	// Output:
	//
}
