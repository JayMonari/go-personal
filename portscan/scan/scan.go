package scan

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const reservedPortMax = 1024
const widePortMax = 49152

type scanResult struct {
	Port  string
	State string
}

// ScanPort scans a single port under a given protocol.
func ScanPort(protocol, hostname string, port int) scanResult {
	sr := scanResult{
		Port:  fmt.Sprintf("%d/%s", port, protocol),
		State: "Closed",
	}
	addr := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		return sr
	}
	defer conn.Close()

	sr.State = "Open"
	return sr
}

// RangeUDP scans a range of ports from the first input to the last.
func RangeUDP(from, to int, hostname string) []scanResult {
	mid := to - from
	ss := make([]scanResult, mid)
	for i := 0; i < mid; i++ {
		ss[i] = ScanPort("udp", hostname, i+from+1)
	}
	return ss
}

// RangeTCP scans a range of ports from the first input to the last.
func RangeTCP(from, to int, hostname string) []scanResult {
	mid := to - from
  wg := sync.WaitGroup{}
  wg.Add(mid)
	ss := make([]scanResult, mid)
	for i := 0; i < mid; i++ {
		go func(i int) {
			ss[i] = ScanPort("tcp", hostname, i+from+1)
      wg.Done()
		}(i)
	}
  wg.Wait()
	return ss
}

// Wide performs a scan from port 1 to 49152 and returns the result of each
// port.
func Wide(hostname string) []scanResult {
	return RangeTCP(0, widePortMax, hostname)
}
