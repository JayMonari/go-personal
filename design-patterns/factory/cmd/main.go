package main

import (
	"factory/factory"
	"fmt"
	"os"
)

func main() {
	var t factory.DBType
	fmt.Print("Type the connection you want 0: Postgres, 1: MySQL")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Exiting from error: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Invalid engine")
		os.Exit(1)
	}

	fmt.Printf("Connection: %v\n", conn)
	// XXX: This will crash without the DBs running.
	// conn.Connect()
	// defer conn.Close()
	// conn.GetNow()
}
