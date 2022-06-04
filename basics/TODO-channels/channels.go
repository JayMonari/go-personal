package channels

import "fmt"

// When using the <- operator, think of blocking
// TODO(jaymonari): explain <- operator in detail in lesson.

// ChannelUnbufferedNoReceiveCausesPanic will cause a panic if you run it by
// itself. This is because there is a very well understood idiom in Go:
//
// "Don't communicate by sharing memory; share memory by communicating."
//   -- Rob Pike
//
// This means that unbuffered channels are **required** to communicate with
// other goroutines, i.e. for every send c <- "stuff" there must be a receive
// <-c from another goroutine.
//
// NOTE(jay): Running this in main will cause a panic!
// fatal error: all goroutines are asleep - deadlock!
func ChannelUnbufferedNoReceiveCausesPanic() {
	c := make(chan string)
	c <- "stuff"
	// Uncommenting this still won't work since it's the same goroutine.
	// wontWork := <-c
}

// ChannelUnbuffered shows how to push values into an unbuffered channel. To be
// unbuffered means there is no limit to the amount of values you can push into
// the channel.
func ChannelUnbuffered(c chan int) {
	for i := 1; i <= 9001; i++ {
		if i == 13 {
			i = 8997
		}
		c <- i
	}
	close(c)
}

// ChannelBuffered shows you that when you buffer a channel that is the final
// capacity that the channel can reach of a certain type and if we try to go
// over we will panic! But we can actually send values into the channel, unlike
// unbuffered channels which would cause a panic.
func ChannelBuffered(c chan string) {
	c <- "first value to channel"
	c <- "second value to channel"
	// NOTE(jay): If we uncomment this line we panic with
	//    fatal error: all goroutines are asleep - deadlock!
	// This is because this buffered channel can only be filled up to
	// its limit -- 2 and since no one is taking the values out, we block
	// indefinitely and find ourselves in a deadlock!
	// c <- "third thing"
	close(c)
}

// ChannelReceiveValueFrom shows you how to declare a channel that will only allow
// values to receieve from the channel.
func ChannelReceiveValueFrom(c <-chan string) {
	// NOTE(jay): This won't work, comiler says:
	//    invalid operation: cannot send to receive-only type <-chan string
	//    [compiler: InvalidSend]
	// c <- ""
	one := <-c
	two := <-c
	three, four := <-c, <-c
	fmt.Println(one, two, three, four)
}

// ChannelSendValueInto shows you how to declare a channel that will only allow
// values to be sent into the channel.
func ChannelSendValueInto(c chan<- string) {
	// NOTE(jay): This won't work, compiler says:
	//    invalid operation: cannot receive from send-only channel c (variable of
	//    type chan<- string) [compiler: InvalidReceive]
	// wontWork := <-c
	c <- "send"
	c <- "all"
	c <- "the"
	c <- "values"
	close(c)
}
