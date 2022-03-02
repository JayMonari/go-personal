package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"syscall"

	"golang.org/x/sys/unix"
)

// XXX: Without a config this happens --
// listen udp 127.0.0.1:2002: bind: address already in use exit status 1

func main() {
	roughtimeCfg := &net.ListenConfig{Control: setListenerOptions}
	roughtimePC, err := roughtimeCfg.ListenPacket(context.Background(), "udp", ":2002")
	must(err)
	go handlePacket("roughtime", roughtimePC)
	spectrumCfg := &net.ListenConfig{Control: setListenerOptions}
	spectrumPC, err := spectrumCfg.ListenPacket(context.Background(), "udp", "127.0.0.1:2002")
	must(err)
	go handlePacket("spectrum", spectrumPC)
	select {}
	// nc -u 127.0.0.1 2002
}

func setListenerOptions(proto, addr string, c syscall.RawConn) error {
	return c.Control(func(descriptor uintptr) {
		syscall.SetsockoptInt(
			int(descriptor),
			unix.SOL_SOCKET,
			// SO_REUSEADDR ensures that roughtime does not block use of port 2002 by
			// Spectrum
			unix.SO_REUSEADDR, 1)
	})
}

func handlePacket(s string, pc net.PacketConn) {
	fmt.Println("WE UP!", pc.LocalAddr().String())
	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			break
		}
		data := buf[:n]
		pc.WriteTo(append([]byte("UDP echo: "), data...), addr)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
