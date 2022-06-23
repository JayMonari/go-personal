package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	go createServer()
	time.Sleep(100 * time.Millisecond)
	go createClient()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT)
	time.Sleep(200 * time.Millisecond)
}

func createClient() {
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, _ := net.DialUDP("udp", nil, addr)
	for i := 0; i < 40; i++ {
		conn.Write([]byte("IDC IF YOU GET THIS WOOOOOOOOOOOOO~~~~" + strconv.Itoa(i)))
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}

func createServer() {
	// Empty host --> INADDR_ANY: listen on host's available unicast+anycast IPs
	pc, err := net.ListenPacket("udp", ":8080") // socket(), bind()
	defer pc.Close()
	// implements net.PacketConn interface
	must(err)
	for {
		buf := make([]byte, 1024)
		size, addr, err := pc.ReadFrom(buf)
		if err != nil {
			break
		}
		data := buf[:size]
		pc.WriteTo(append([]byte("UDP echo: "), data...), addr)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
