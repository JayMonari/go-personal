package main

import (
	"fmt"
	"portscan/scan"
)

func main() {
	fmt.Println("Port Scanning")
	results := scan.Wide("localhost")
	fmt.Println(results)
}
