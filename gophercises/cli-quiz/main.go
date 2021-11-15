package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("Failed to open the CSV file: %s\n", *csvFilename)
	}
	r := csv.NewReader(file)
	rr, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read the CSV file: %v\n", err)
	}
	bufRdr := bufio.NewReader(os.Stdin)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, r := range parseRecords(rr) {
		fmt.Printf("Problem #%d: %s = ", i+1, r.ques)
		ansCh := make(chan string, 1)
		go func() {
			ans, err := bufRdr.ReadString('\n')
			if err != nil {
				log.Fatalf("Failed to read your answer: %v", err)
			}
			ansCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(rr))
			return
		case ans := <-ansCh:
			if ans[:len(ans)-1] == r.ans {
				correct++
				fmt.Printf("YOU MADE IT! %s", ans)
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(rr))
}

func parseRecords(rr [][]string) []problem {
	ret := make([]problem, len(rr))
	for i, r := range rr {
		ret[i] = problem{
			ques: r[0],
			ans:  strings.TrimSpace(r[1]),
		}
	}
	return ret
}

type problem struct {
	ques, ans string
}
