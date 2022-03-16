package main

import (
	"pass"

	"github.com/atotto/clipboard"
)

func main() {
	p := pass.Gen(16)

	clipboard.WriteAll(p)
}
