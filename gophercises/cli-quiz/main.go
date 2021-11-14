package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
  csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
  flag.Parse()

  file, err := os.Open(*csvFilename)
  if err != nil {
    log.Fatalf("Failed to open the CSV file: %s\n", *csvFilename)
  }
  fmt.Println(file)
}
