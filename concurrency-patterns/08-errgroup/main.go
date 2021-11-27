package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	// wait := waitGroups()
	wait := errGroup()

	<-wait
}

func errGroup() <-chan struct{} {
	ch := make(chan struct{}, 1)

	var g errgroup.Group

	for _, name := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		// Required because of closure and how goroutines will work to grab the
		// same value at back of range loop
		n := name

		g.Go(func() error {
			ch, err := read(n)
			if err != nil {
				return fmt.Errorf("error reading %w", err)
			}

			for line := range ch {
				fmt.Println(n, line)
			}
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading files %v", err)
		}
		close(ch)
	}()

	return ch
}

func waitGroups() chan struct{} {
	ch := make(chan struct{}, 1)

	var wg sync.WaitGroup

	for _, name := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			ch, err := read(n)
			if err != nil {
				panic("bad file name")
			}

			for line := range ch {
				fmt.Println(n, line)
			}
		}(name)
	}

	go func() {
		wg.Wait()

		close(ch)
	}()

	return ch
}

func read(name string) (<-chan []string, error) {
	f, err := os.Open(name)
	if err != nil {
		panic("bad file name")
	}

	ch := make(chan []string)
	go func() {
		cr := csv.NewReader(f)
		for {
			r, err := cr.Read()
			if errors.Is(err, io.EOF) {
				close(ch)
				return
			}
			ch <- r
		}
	}()

	return ch, nil
}
