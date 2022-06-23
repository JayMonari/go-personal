package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go createServer()
	time.Sleep(100 * time.Millisecond)
	go createClient()
	time.Sleep(300 * time.Millisecond)
}

func createServer() {
	// Empty host --> INADDR_ANY: listen on host's available unicast+anycast IPs
	ln, err := net.Listen("tcp", ":8080") // socket(), bind(), listen()
	// implements net.Conn interface
	defer ln.Close()
	check(err)
	log.Print(ln.Addr().String())
	log.Printf("Started listening (%s) on port %s", "tcp", "8080")
	conn, err := ln.Accept()
	defer conn.Close()
	check(err)
	for {
		fmt.Println("we made it")
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		fmt.Println("buf:", string(buf[:size]))
		if err != nil {
			break
		}
		data := buf[:size]
		conn.Write(append([]byte("TCP echo: "), data...))
	}
}

func createClient() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8080")
	conn, _ := net.DialTCP("tcp", nil, addr)
	for i := 0; i < 100; i++ {
		conn.Write([]byte("GOOD DAY SIR!!!"))
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
