package main

import (
	"adapter"
	"fmt"
)

func main() {
	var t string
	fmt.Println("Choose your type of transportation (bike, car):")
	fmt.Scanln(&t)

  tran := adapter.Factory(t)
  tran.Move()
}
