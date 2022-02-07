package main

import (
	"fmt"
	"time"
)

func main() {
	processData("main goroutine")
	go processData("new goroutine1")
	go processData("new goroutine2")
	go processData("new goroutine3")

	go func(comingFrom string) {
		fmt.Println("coming from:", comingFrom)
	}("anonymous inline goroutine")

	go processData("new goroutine4")
	go processData("new goroutine5")
	go processData("new goroutine6")

	// We have to wait, because the main goroutine will shutdown other goroutines
	// and exit immediately. Comment out this line and the "time" package import
	// above and see what you get!
	time.Sleep(time.Second)
	fmt.Println("exiting main")
}

func processData(comingFrom string) {
	fmt.Println("coming from:", comingFrom)
}
