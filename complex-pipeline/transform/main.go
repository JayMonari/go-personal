package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open(parseFlags())
	if err != nil {
		log.Fatalf("Opening file %v", err)
	}
	defer f.Close()

	scr := bufio.NewScanner(f)
	for scr.Scan() {
		fmt.Printf("%#v\n", strings.Split(strings.Trim(scr.Text(), "\n"), "\t"))
	}
}

func parseFlags() string {
	var file string
	flag.StringVar(&file, "file", "name.basics.tsv", "tab separated values file")
	flag.Parse()
	return file
}
