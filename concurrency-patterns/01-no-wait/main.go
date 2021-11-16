package main

import "fmt"

func main() {
	go hello()
}

func hello() {
	fmt.Print(`
Never going to see this because when main exits, all goroutines are stopped.
This is due to the concurrent nature of goroutines.
`)
}
