package channels_test

import (
	"basics/channels"
	"fmt"
	"time"
)

func ExampleChannelUnbufferedNoReceiveCausesPanic() {
	go channels.ChannelUnbufferedNoReceiveCausesPanic()
	// Output:
}

func ExampleChannelUnbuffered() {
	c := make(chan int)
	done := make(chan struct{}, 1)

	go channels.ChannelUnbuffered(c)

	go func(c chan int, done chan struct{}) {
		fmt.Println("OVER 9000!!!")
		for n := range c {
			fmt.Print(n, " ")
		}
		done <- struct{}{}
	}(c, done)

	<-done

	// Output:
	// OVER 9000!!!
	// 1 2 3 4 5 6 7 8 9 10 11 12 8997 8998 8999 9000 9001
}

func ExampleChannelBuffered() {
	c := make(chan string, 2)
	done := make(chan struct{}, 1)

	go channels.ChannelBuffered(c)

	go func(c chan string, done chan struct{}) {
		for s := range c {
			fmt.Println(s)
		}
		done <- struct{}{}
	}(c, done)

	<-done

	// Output:
	// first value to channel
	// second value to channel
}

func ExampleChannelSendValueInto() {
	c := make(chan string)

	go channels.ChannelSendValueInto(c)

	go channels.ChannelReceiveValueFrom(c)

	time.Sleep(5 * time.Millisecond)

	// Output:
	// send all the values
}

func ExampleChannelReceiveValueFrom() {
	c := make(chan string)

	go channels.ChannelSendValueInto(c)

	go channels.ChannelReceiveValueFrom(c)

	time.Sleep(5 * time.Millisecond)

	// Output:
	// send all the values
}
