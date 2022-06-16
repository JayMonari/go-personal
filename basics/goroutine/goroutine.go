package goroutine

import (
	"fmt"
	"time"
)

// https://go.dev/talks/2012/concurrency.slide#53

func WillNotWait() {
	// NOTE(jay): This will be seen if we run `go test` for `ExampleWillNotWait`
	// it just won't be a part of the main goroutines output because it exits.
	go toofast()
}

func toofast() { fmt.Println("We'll never see this... without waiting") }

// SwitchToOther shows us how to artificially allow the goroutine we spawn to
// finish and exit, by slowing down the main goroutine. This is **not** how
// it's done in go. We use channels for true concurrency, but this is important
// to see before we introduce channels.
func SwitchToOther() {
	go toofast()
	// Make it wait 8 milliseconds to see separate goroutines output.
	time.Sleep(8 * time.Millisecond)
}

func AnonymousFunctions(val any) {
	// The `go` keyword needs a function and that is all, even if it is an
	// anonymous function, it can still be used in a goroutine
	go func(comingFrom string) {
		fmt.Println("coming from:", comingFrom)
	}("anonymous function goroutine")

	go func(v any) {
		switch t := val.(type) {
		case string:
			fmt.Printf("you chose %T: %s\n", t, t)
		case int:
			fmt.Printf("you chose %T: %d\n", t, t)
		case bool:
			fmt.Printf("you chose %T: %t\n", t, t)
		case float64:
			fmt.Printf("you chose %T: %f\n", t, t)
		case []struct{}:
			fmt.Printf("you chose %T: %#v\n", t, t)
		default:
			fmt.Printf("what is this 👀 %T: %#v\n", t, t)
		}
	}(val)

	// NOTE(jay): We have to wait (`time.Sleep`), because the main goroutine will
	// shutdown other goroutines and exit immediately. Comment out this line and
	// see what you get!
	time.Sleep(8 * time.Millisecond)
	fmt.Println("👋👋👋 Time to exit")
}

// NoOrder shows that asynchronous truly means there is no determined order.
// That the goroutines are not in sync and are not in serialized order. We
// spawn several goroutines and without sorting the outputs **this output will
// never be deterministic** this means the order is determined by which
// goroutine got time scheduled first.
func NoOrder() {
	for i := 0; i < 3; i++ {
		go processData(fmt.Sprintf("goroutine%d", i))
	}
	go processData("goroutine3")
	go processData("goroutine4")
	go processData("goroutine5")
}

func processData(comingFrom string) {
	fmt.Println("coming from:", comingFrom)
}
