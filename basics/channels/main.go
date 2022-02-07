package main

import "fmt"

// When using the <- operator, think of blocking
func main() {
	messages := make(chan string)
	go ChannelUnbuffered(messages)
	blockUntilRecieve := <-messages
	fmt.Println("received from channel:", blockUntilRecieve)
	messages = ChannelBuffered()
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// XXX: Running this in main will cause a panic
// fatal error: all goroutines are asleep - deadlock!
func ChannelCausesPanic() {
	msgs := make(chan string)
	msgs <- "stuff"
	fmt.Println("msgs:", msgs)
}

func ChannelUnbuffered(msgs chan string) {
	msgs <- "stuff"
	fmt.Println("msgs:", msgs)
}

func ChannelBuffered() chan string {
	msgs := make(chan string, 2)
	msgs <- "first thing"
	msgs <- "second thing"
	// XXX: If we uncomment this line we panic with
	//   fatal error: all goroutines are asleep - deadlock!
	// This is because a buffered channel can only be filled up to its limit - 2
	// since no one is taking the values out, we block indefinitely.
	// msgs <- "third thing"
	return msgs
}
