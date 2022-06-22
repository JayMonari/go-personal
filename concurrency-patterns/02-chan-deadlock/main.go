package main

import "fmt"

// XXX(jay): This program panics because there is no goroutine outside of
// `main` interacting with the `ch` channel:
//
// fatal error: all goroutines are asleep - deadlock!
func main() {
	ch := make(chan int)
	ch <- 10
	v := <-ch
	fmt.Println("recieved", v)
}
