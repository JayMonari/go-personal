package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v", err))
	}

	br1 := breakup("1", ch1)
	br2 := breakup("2", ch1)
	br3 := breakup("3", ch1)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		// Dealing with channels that have been closed is a great use for a select
		// statement.
		select {
		case _, open := <-br1:
			if !open {
				br1 = nil
			}
		case _, open := <-br2:
			if !open {
				br2 = nil
			}
		case _, open := <-br3:
			if !open {
				br3 = nil
			}
		}
	}

	fmt.Println("jobs done!")
}

func read(name string) (<-chan []string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	cr := csv.NewReader(f)

	ch := make(chan []string)

	// Why do we need this in a separate goroutine? Well the application will
	// deadlock otherwise...
	go func() {
		for {
			r, err := cr.Read()
			if err == io.EOF {
				close(ch)
				break
			}
			ch <- r
		}
	}()

	return ch, nil
}

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}
		close(chE)
	}()

	return chE
}
