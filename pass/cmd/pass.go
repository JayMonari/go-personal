package main

import (
	"os/exec"
	"pass"

	"github.com/atotto/clipboard"
)

func main() {
	p := pass.Gen(16)

	cmd := exec.Command("xsel", "-ib")
	clipboard.WriteAll(p)
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	in.Write([]byte(p))
}
