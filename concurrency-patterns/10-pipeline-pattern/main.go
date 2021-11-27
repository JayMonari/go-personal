package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	recCh, err := read("file1.csv")
	if err != nil {
		log.Fatalf("Could not read csv %v", err)
	}

	for val := range titeleize(sanitize(recCh)) {
		fmt.Printf("%v\n", val)
	}
}

func read(name string) (<-chan []string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	ch := make(chan []string)
	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3
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

func sanitize(strCh <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strCh {
			if len(val[0]) > 3 {
				fmt.Println("skipped ", val)
				continue
			}
			ch <- val
		}
		close(ch)
	}()

	return ch
}

// titeleize uppercases the first value and swap the next two... because that's
// what the PM wants.
func titeleize(strCh <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strCh {
			val[0] = strings.Title(val[0])
			val[1], val[2] = val[2], val[1]

			ch <- val
		}
		close(ch)
	}()

	return ch
}
