package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	scheme = "http"
	host   = "localhost"
	port   = 9001
)

var (
	uri = fmt.Sprintf("%s://%s:%d", scheme, host, port)
)

func main() {
	action, arg1 := os.Args[1], os.Args[2]
	curl := exec.Command(
		"curl", uri,
		"--data", "FIXME",
	)
	switch action {
	case "login":
		curl.Args[1] += "/users/login"
		curl.Args[3] = fmt.Sprintf(`{"username":"%s","password":"password123"}`, arg1)
		fmt.Println(curl.Args)
		out, err := curl.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		if err := curl.Start(); err != nil {
			log.Fatal(err)
		}

		jq := exec.Command("jq")
		jq.Stdin, jq.Stdout, jq.Stderr = out, os.Stdout, os.Stderr
		if err := jq.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
