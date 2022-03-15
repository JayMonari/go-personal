package keywords

// package
// import

// func
// go
// defer
// return

// for
// range
// continue
// break

// var/const

// MyType shows the three different ways that the `type` keyword is used.
// 1. Creating a new `type`
// 2. Creating a new `struct`
// 3. Creating a new `interface`
func MyType() {
	// Creating new types can have many benefits, such as stricter typing,
	// clearer APIs, and custom methods.
	type myNewType int

	// `struct` is short for data structure. It is named this way because it is a
	// place to structure or frame or collect similar data that has something in
	// common.
	type myStruct struct {
		ExportedField   int
		unexportedField string
	}

	// Often the name of an interface comes from the method it defines plus "-er"
	// e.g. from the standard library: fmt.Stringer, io.Reader, io.Writer
	// other examples: Printer, Greeter, Walker, Swimmer, Waiter, Sender
	type myHeater interface {
		Heat() string
	}
}

// map

// goto

// chan
// select

// if | else if | else

// switch
// case
// fallthrough
// default
