package main

import (
	"fmt"
	"reflect"
)

// NOTE(docs): See "The Laws of Reflection" for an introduction to reflection in Go:
// https://golang.org/doc/articles/laws_of_reflection.html

// NOTE(jay): tl;dr of the Laws:
// Interfaces are King/Queen 👑 in Go. They have two ✌️  things inside of them; a **VALUE**
// and a **TYPE** e.g. interface(VALUE, concrete TYPE) ➡️a interface(myVar, *http.Request)
// or interface(myBigInt, *big.Int) because of these properties interfaces are very useful
// for getting information at runtime.
//
// The only entrance points are [reflect.TypeOf] and [reflect.ValueOf], which both take
// [any] (the empty interface) as their arguement. I would suggest starting there if
// you've never used this package before.
//
// The 3 laws:
//   1. Reflection goes from (interface package of 2) to reflection object.
//   2. Reflection goes from reflection object to (interface package of 2).
//   3. To modify a reflection object, the value must be settable.

const (
	Ptr = reflect.Pointer

	Invalid reflect.Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

const (
	_             reflect.SelectDir = iota
	SelectSend                      // case Chan <- Send
	SelectRecv                      // case <-Chan:
	SelectDefault                   // default
)

const (
	RecvDir reflect.ChanDir     = 1 << iota // <-chan
	SendDir                                 // chan<-
	BothDir = RecvDir | SendDir             // chan
)

// func ExampleCopy() {
// 	reflect.Copy()
// 	// Output:
// }
//
// func ExampleDeepEqual() {
// 	reflect.DeepEqual()
// 	// Output:
// }
//
// func ExampleSelect() {
// 	reflect.Select()
// 	// Output:
// }
//
// func ExampleSwapper() {
// 	reflect.Swapper()
// 	// Output:
// }
//
// // type ChanDir int
// // const RecvDir ChanDir = 1 << iota ...
// // type Kind uint
// // const Invalid Kind = iota ...
// // type MapIter struct{ ... }
// // type Method struct{ ... }
// // type SelectCase struct{ ... }
// // type SelectDir int
// // const SelectSend SelectDir ...
// // type SliceHeader struct{ ... }
// // type StringHeader struct{ ... }
// // type StructField struct{ ... }
// func ExampleVisibleFields() {
// 	reflect.VisibleFields()
// 	// Output:
// }
//
// // type StructTag string
// // type Type interface{ ... }
// func ExampleArrayOf() {
// 	reflect.ArrayOf()
// 	// Output:
// }
//
// func ExampleChanOf() {
// 	reflect.ChanOf()
// 	// Output:
// }
//
// func ExampleFuncOf() {
// 	reflect.FuncOf()
// 	// Output:
// }
//
// func ExampleMapOf() {
// 	reflect.MapOf()
// 	// Output:
// }
//
// func ExamplePointerTo() {
// 	reflect.PointerTo()
// 	// Output:
// }
//
// func ExamplePtrTo() {
// 	reflect.PtrTo()
// 	// Output:
// }
//
// func ExampleSliceOf() {
// 	reflect.SliceOf()
// 	// Output:
// }
//
// func ExampleStructOf() {
// 	reflect.StructOf()
// 	// Output:
// }
//
// func ExampleTypeOf() {
// 	reflect.TypeOf()
// Output:
// }
//
// // type Value struct{ ... }
// func ExampleAppend() {
// 	reflect.Append()
// 	// Output:
// }
//
// func ExampleAppendSlice() {
// 	reflect.AppendSlice()
// 	// Output:
// }
//
// func ExampleIndirect() {
// 	reflect.Indirect()
// 	// Output:
// }
//
// func ExampleMakeChan() {
// 	reflect.MakeChan()
// 	// Output:
// }
//
// func ExampleMakeFunc() {
// 	reflect.MakeFunc()
// Output:
// }
//
// func ExampleMakeMap() {
// 	reflect.MakeMap()
// 	// Output:
// }
//
// func ExampleMakeMapWithSize() {
// 	reflect.MakeMapWithSize()
// 	// Output:
// }
//
// func ExampleMakeSlice() {
// 	reflect.MakeSlice()
// 	// Output:
// }
//
// func ExampleNew() {
// 	reflect.New()
// 	// Output:
// }
//
// func ExampleNewAt() {
// 	reflect.NewAt()
// 	// Output:
// }

func ExampleValueOf() {
	// Uintptr
	// Complex64
	// Complex128
	// Array
	// Chan
	// Func
	// Interface
	// Map
	// Struct
	// UnsafePointer
	fmt.Printf("%q\n", reflect.ValueOf(nil).String())
	fmt.Printf("%q\n", reflect.ValueOf(true).String())
	fmt.Printf("%q\n", reflect.ValueOf(int16(0)).String())
	fmt.Printf("%q\n", reflect.ValueOf(uint32(0)).String())
	fmt.Printf("%q\n", reflect.ValueOf(2.5).String())
	fmt.Printf("%q\n", reflect.ValueOf(&struct{ pointer *int }{}).String())
	fmt.Printf("%q\n", reflect.ValueOf("STRING").String())
	fmt.Printf("%q\n", reflect.ValueOf([]byte("slice")).String())
	// Output:
}

// func ExampleZero() {
// 	reflect.Zero()
// 	// Output:
// }
//
// // type ValueError struct{ ... }
