package channel

import (
	"fmt"
	"time"
)

// KidEats imitates a goroutine (this function in this case) taking in a
// resource, doing something with it and then returning some result.
func KidEats(candy string) string {
	fmt.Println("🧒 💬 I'm eating " + candy)
	return "wrapper"
}

// UnbufferedNoReceiveCausesPanic will cause a panic if you run it by
// itself. This is because there is a very well understood idiom in Go:
//
// "Don't communicate by sharing memory; share memory by communicating."
//
//	-- Rob Pike
//
// This means that unbuffered channels are **required** to communicate with
// other goroutines, i.e.
// for every send `c <- "stuff"`
// there must be a receive `giveme := <-c` from another goroutine.
//
// XXX(jay): Running this in main will cause a panic!
//
//	fatal error: all goroutines are asleep - deadlock!
func UnbufferedNoReceiveCausesPanic() {
	c := make(chan string)
	c <- "doesn't matter going to panic"
	fmt.Println("This doesn't work", <-c, "there is no communication")
}

// Unbuffered shows how to send values into an unbuffered channel. To be
// unbuffered means there is no limit to the amount of values you can push into
// the channel.
func Unbuffered(c chan int) {
	for i := 1; i <= 9001; i++ {
		c <- i
	}
	close(c)
}

// BufferedChanWorks demonstrates you can use a buffered channel in the same
// goroutine without creating a deadlock because there is an area to store the
// extra values that will be pushed into the channel, a buffer.
func BufferedChanWorks() {
	c := make(chan string, 1)
	c <- "This is a useless to do, but great for understanding"
	// NOTE(jay): We could not do this
	//   c <- "Some other value into the buffed channel with only 1 opening"
	// because it _would_ block waiting for space to put the value in.
	fmt.Println(<-c)
}

// Buffered shows you that when you buffer a channel that is the final
// capacity that the channel can reach of a certain type and if we try to go
// over we will panic! But we can actually send values into the channel, unlike
// unbuffered channels which would cause a panic.
func Buffered(c chan string) {
	c <- "first value to channel"
	c <- "second value to channel"
	// XXX(jay): If we uncomment this line we panic with
	//    fatal error: all goroutines are asleep - deadlock!
	// This is because this buffered channel can only be filled up to
	// its limit -- 2 and since no one is taking the values out, we block
	// indefinitely and find ourselves in a deadlock!
	// c <- "third thing"
	close(c)
}

// CommunicateDone is an example of how to let a completely separate goroutine
// (which has no idea what any other goroutines are doing) know when it is safe
// to continue execution by communicating it has finished the work that it was
// tasked with doing.
func CommunicateDone(done chan<- struct{}) {
	fmt.Println("Starting up to do some hard work.")
	time.Sleep(250 * time.Millisecond)
	fmt.Println("All finished with the hard work.")
	done <- struct{}{}
}

// ClosingChannels is an example function that shows there is a second value you can look
// at when receiving a value from a channel. This second value tells you if the channel is
// still open, which is a good clue (if the value is the zero value of its type)
// that there are no more values to receive on that channel.
func ClosingChannels(kitchen <-chan rune, done chan<- struct{}) {
	for {
		order, open := <-kitchen
		if !open {
			fmt.Println("channel has closed; zero value of channel:", order)
			done <- struct{}{}
			return
		}
		fmt.Printf("requested to make: %q\n", order)
	}
}

// Range shows how to range over both a buffered and unbuffered channel. It is good to
// note that neither of the channels are closed by this function. It's almost never a good
// idea for the receiver to be closing the channels.
func Range(buf, unbuf chan string, done chan struct{}) {
	go func() {
		for s := range buf {
			fmt.Println(s)
		}
		done <- struct{}{}
	}()
	go func() {
		for s := range unbuf {
			fmt.Println(s)
		}
		done <- struct{}{}
	}()
}

// ReceiveValueFrom shows you how to declare a channel that will only
// allow values to receieve from the channel.
func ReceiveValueFrom(c <-chan string) {
	// XXX(jay): This won't work, compiler says:
	//    invalid operation: cannot send to receive-only type <-chan string
	//    [compiler: InvalidSend]
	// c <- ""
	one := <-c
	two := <-c
	three, four := <-c, <-c
	fmt.Println(one, two, three, four)
}

// SendValueInto shows you how to declare a channel that will only allow
// values to be sent into the channel.
func SendValueInto(c chan<- string) {
	// XXX(jay): This won't work, compiler says:
	//  invalid operation: cannot receive from send-only channel c (variable of
	//  type chan<- string) [compiler: InvalidReceive]
	// wontWork := <-c
	c <- "send"
	c <- "all"
	c <- "the"
	c <- "values"
	close(c)
}

// Select shows off all of the features of the [select] keyword. It's very similar to the
// [switch] keyword, but for concurrency. In here we see that the best features of
// [select] are it's ability to make all channel types non-blocking on sends and receives
// with the [default] keyword and also the ability to create multi-branching paths.
func Select(fullbuf, unbuf chan string, done chan struct{}) {
	select {
	case fullbuf <- "cannot push this into a full buffer":
	default:
		fmt.Println("Did not block after trying to send value into a full buffered channel.")
	}

	select {
	case recvVal := <-unbuf:
		fmt.Println("We will never receive this value because no one is sending", recvVal)
	case <-unbuf:
		// NOTE(jay): We don't have to capture the value, maybe we just want to know if there
		// are values in the buffer. Again, this would block under normal circumstances.
	default:
		fmt.Println("Did not block after trying to receive value from an unbuffered channel")
	}
	done <- struct{}{}

	for {
		select {
		case v := <-fullbuf:
			fmt.Println(v)
		case v := <-unbuf:
			fmt.Println(v)
		case <-time.After(100 * time.Millisecond): // Just like a timeout.
			done <- struct{}{}
			return
			// NOTE(jay): In this case we would get into trouble if we used the [default] case
			// because it's NOT verifiable that after the `case v := <-unbuf` the next value for
			// `unbuf <- "some string"` will be sent before the next iteration in this for loop.
			// case default:
			// done <- struct{}{}
			// return
		}
	}
}
