package main

import (
	"fmt"
	"proxy"
	"time"
)

var loc proxy.Proxy

func main() {
  loc = proxy.NewLocal()

  TimeByID(2)
  TimeByID(2)
  TimeByID(1)
  TimeByID(2)
  TimeByID(3)
  TimeByID(5)
  TimeByID(4)
  TimeAll()
}

func TimeByID(ID uint) {
  start := time.Now()
  fmt.Printf("%+v", loc.GetByID(ID))
  elapsed := time.Since(start)
  fmt.Printf("Time elapsed: %v\n", elapsed)
}

func TimeAll() {
  start := time.Now()
  fmt.Printf("%+v", loc.GetAll())
  elapsed := time.Since(start)
  fmt.Printf("Time elapsed: %v\n", elapsed)
}
