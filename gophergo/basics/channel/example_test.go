package channel_test

import (
	"fmt"
	"time"

	"basics/channel"
)

func ExampleKidEats() {
	teacher := make(chan string)
	// This function could easily be called HandoutCandy
	go func() {
		teacher <- "🍭"
		teacher <- "🍬"
		teacher <- "🍫"
		teacher <- "🧁"
		teacher <- "🍪"
	}()

	garbage := make(chan string)
	for i := 0; i < 5; i++ {
		go func() {
			garbage <- channel.KidEats(<-teacher)
		}()
	}

	bin := make([]string, 5)
	for i := 0; i < 5; i++ {
		// instead of throwing away the result, we would use it in a real app.
		bin[i] = <-garbage
	}
	fmt.Println("All the wrappers have been collected.")
	// Unordered output:
	// 🧒 💬 I'm eating 🍭
	// 🧒 💬 I'm eating 🍬
	// 🧒 💬 I'm eating 🍪
	// 🧒 💬 I'm eating 🍫
	// 🧒 💬 I'm eating 🧁
	// All the wrappers have been collected.
}

func ExampleUnbufferedNoReceiveCausesPanic() {
	// XXX(jay): Removing this from a separate goroutine will block forever.
	go channel.UnbufferedNoReceiveCausesPanic()
	// Output:
}

func ExampleUnbuffered_drain_channel() {
	c := make(chan int)
	done := make(chan struct{}, 1)

	go channel.Unbuffered(c)

	go func(c chan int, done chan struct{}) {
		for n := range c {
			if n > 12 && n < 8998 { // too much output to care about
				continue
			}
			if n > 9000 {
				fmt.Print(n)
				break
			}
			fmt.Print(n, " ")
		}
		done <- struct{}{}
	}(c, done)

	// NOTE(jay): This is how we make the main goroutine,
	// `ExampleUnbuffered_drain_channel`, wait for the 2 other goroutines we spun
	// up. If we remove this then it would exit without letting the others do
	// their job.
	<-done

	fmt.Println("\nOVER 9000!!!")

	// Output:
	// 1 2 3 4 5 6 7 8 9 10 11 12 8998 8999 9000 9001
	// OVER 9000!!!
}

func ExampleUnbuffered_channel_has_values_left() {
	c := make(chan int)
	go channel.Unbuffered(c)
	<-c // drain values because we can
	fmt.Println(<-c)
	// drain value after to show we don't block because we have more values to
	// receive.
	<-c
	// Output: 2
}

func ExampleBufferedChanWorks() {
	channel.BufferedChanWorks()
	// Output: This is a useless to do, but great for understanding
}

func ExampleBuffered() {
	c := make(chan string, 2)
	done := make(chan struct{}, 1)

	go channel.Buffered(c)

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

func ExampleCommunicateDone() {
	done := make(chan struct{}, 1)
	go channel.CommunicateDone(done)
	<-done
	// Output:
	// Starting up to do some hard work.
	// All finished with the hard work.
}

func ExampleClosingChannels() {
	done := make(chan struct{})
	kitchen := make(chan rune)

	go channel.ClosingChannels(kitchen, done)
	for _, order := range []rune{'🍔', '🍲', '🍗', '🍟', '🥞', '🍕', '🥪'} {
		kitchen <- order
	}
	close(kitchen)
	fmt.Println("All orders have been sent to the kitchen.")

	// The output (for orders) is **always** ordered, do you know why? 🤔 Think 💪
	<-done
	// Unordered output:
	// requested to make: '🍔'
	// requested to make: '🍲'
	// requested to make: '🍗'
	// requested to make: '🍟'
	// requested to make: '🥞'
	// requested to make: '🍕'
	// requested to make: '🥪'
	// All orders have been sent to the kitchen.
	// channel has closed; zero value of channel: 0
}

func ExampleRange() {
	buf := make(chan string, 4)
	unbuf := make(chan string)
	done := make(chan struct{})

	buf <- "It's possible to close"
	buf <- "a non-empty channel and"
	buf <- "still have the remaining"
	buf <- "values be received"
	close(buf)
	go channel.Range(buf, unbuf, done)

	unbuf <- "Unbuffered channels"
	unbuf <- "will block on a v := <-receive"
	unbuf <- "and on a send <- v"
	unbuf <- "this is important to remember"
	close(unbuf)

	<-done
	<-done
	// Unordered output:
	// Unbuffered channels
	// will block on a v := <-receive
	// and on a send <- v
	// this is important to remember
	// It's possible to close
	// a non-empty channel and
	// still have the remaining
	// values be received
}

func ExampleSendValueInto() {
	c := make(chan string)

	go channel.SendValueInto(c)

	go channel.ReceiveValueFrom(c)

	// XXX(jay): This is lazy 🦥 should use a `done` channel in its place.
	time.Sleep(5 * time.Millisecond)
	// Output: send all the values
}

func ExampleReceiveValueFrom() {
	c := make(chan string, 4)

	c <- "not necessary to"
	c <- "close a channel"
	c <- "if the other goroutine"
	c <- "isn't draining it"

	go channel.ReceiveValueFrom(c)

	// XXX(jay): This is lazy 🦥 should use a `done` channel in its place.
	time.Sleep(5 * time.Millisecond)
	// Output: not necessary to close a channel if the other goroutine isn't draining it
}

func ExampleSelect() {
	fullbuf := make(chan string, 5)
	unbuf := make(chan string)
	done := make(chan struct{})

	fullbuf <- "A [select] does not guarantee"
	fullbuf <- "execution of case ordering"
	fullbuf <- "meaning the second case may be"
	fullbuf <- "selected to fire even if the"
	fullbuf <- "first case is ready to fire"

	go channel.Select(fullbuf, unbuf, done)
	<-done
	unbuf <- "If one or more of the communications"
	unbuf <- "can proceed, a single one that can"
	unbuf <- "proceed is chosen via a uniform"
	unbuf <- "pseudo-random selection."
	<-done

	// Unordered output:
	// Did not block after trying to send value into a full buffered channel.
	// Did not block after trying to receive value from an unbuffered channel
	// A [select] does not guarantee
	// execution of case ordering
	// meaning the second case may be
	// selected to fire even if the
	// first case is ready to fire
	// If one or more of the communications
	// can proceed, a single one that can
	// proceed is chosen via a uniform
	// pseudo-random selection.
}
