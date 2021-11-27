package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v", err))
	}

	ch2, err := read("file2.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file2 %v", err))
	}

	exit := make(chan struct{})

	chM := merge2(ch1, ch2)

	go func() {
		for v := range chM {
			fmt.Println(v)
		}

		close(exit)
	}()

	<-exit

	fmt.Println("jobs done!")
}

func read(name string) (<-chan []string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			r, err := cr.Read()
			if err == io.EOF {
				close(ch)
				return
			}
			ch <- r
		}
	}()

	return ch, nil
}

func merge1(cc ...<-chan []string) <-chan []string {
	var wg sync.WaitGroup
	wg.Add(len(cc))

	out := make(chan []string)

	send := func(c <-chan []string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cc {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// merge2 works exactly the same as merge1 except we implement our own
// WaitGroup without the use of the sync package. This is done by creating a
// channel that has the length of the amount of passed in channels and upon
// merging their complete contents into the out <-chan a token is placed into
// the wait channel which signals that we can decrement the count by one, just
// like a WaitGroup.
func merge2(cc ...<-chan []string) <-chan []string {
	semaphore := len(cc)
	wait := make(chan struct{}, semaphore)

	out := make(chan []string)

	send := func(c <-chan []string) {
		defer func() { wait <- struct{}{} }()

		for n := range c {
			out <- n
		}
	}

	for _, c := range cc {
		go send(c)
	}

	go func() {
		for range wait {
			semaphore--
			if semaphore == 0 {
				break
			}
		}
		close(out)
	}()

	return out
}
