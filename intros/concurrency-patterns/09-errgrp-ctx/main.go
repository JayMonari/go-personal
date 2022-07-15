package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// ctx := context.Background()
  ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
  defer cancel()
	wait := errGroup(ctx)
	<-wait
}

// errGroup handles parallel requests with a specific timeout; propogate
// cancellation details or messages across multiple goroutines.
func errGroup(ctx context.Context) (exit chan struct{}) {
  exit = make(chan struct{}, 1)
	g, ctx := errgroup.WithContext(ctx)

	for _, name := range []string{"file1.csv", "file2.csv", "file3.csv"} {
		n := name

		g.Go(func() error {
			ch, err := read(n)
			if err != nil {
				return fmt.Errorf("error reading %w", err)
			}

			for {
				select {
				case <-ctx.Done():
					fmt.Println("context completed:", ctx.Err())
					return ctx.Err()
				case line, open := <-ch:
					if !open {
						return nil
					}
					fmt.Println(line)
				}
			}
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("Error reading files: %v", err)
		}

		close(exit)
	}()

	return exit
}

func read(name string) (<-chan []string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	ch := make(chan []string)
	go func() {
		cr := csv.NewReader(f)
		// XXX: Intentional sleep
		time.Sleep(time.Millisecond)
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
