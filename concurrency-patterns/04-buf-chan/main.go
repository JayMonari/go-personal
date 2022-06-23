package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
		}

		// XXX(jay): There could be cases where this message is not completed,
		// this is known as a race condition as the main function exits before this
		// line in the goroutine can be evaluated.
		fmt.Println(time.Now(), "all complete")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "recieved", <-ch)
	fmt.Println(time.Now(), "recieved", <-ch)
	fmt.Println(time.Now(), "recieved", <-ch)

	fmt.Println(time.Now(), "exiting")
}
