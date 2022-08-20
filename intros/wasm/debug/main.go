package main

import "fmt"

func main() {
	fmt.Println("Go WASM")
}

//export multiply
func multiply(a, b int) int { return a * b }
