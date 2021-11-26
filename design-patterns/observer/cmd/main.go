package main

import (
	"bufio"
	"fmt"
	"log"
	"observer"
	"observer/email"
	"observer/message"
	"observer/slack"
	"os"
)

func main() {
	m := message.Message{}
	for readAddMore() {
		obsr := readObserver()
		m.AddObserver(obsr, observerFactory(obsr))
	}

	m.Data = readMessage()
	m.NotifyObservers()
}

func readMessage() string {
	fmt.Print("Type your message ")
	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("enable to read what was typed in: %v", err)
	}
	return text[:len(text)-1]
}

func readObserver() string {
	fmt.Print("Which observer do you want? ")
	r := bufio.NewReader(os.Stdin)
	text, err := r.ReadString('\n')
	if err != nil {
		log.Fatalf("enable to read what was typed in: %v", err)
	}
	return text[:len(text)-1]
}

func readAddMore() bool {
	fmt.Print("Do you want to get more observers? [y/n] ")
	r := bufio.NewReader(os.Stdin)

	if char, _, err := r.ReadRune(); err != nil {
		log.Fatalf("enable to read what was typed in: %v", err)
	} else if char == 'y' {
		return true
	}

	return false
}

func observerFactory(name string) observer.Observer {
	switch name {
	case "slack":
		return &slack.Slack{}
	case "email":
		return &email.Email{}
	default:
		panic(fmt.Sprintf("No product found for that observer: %s", name))
	}
}
