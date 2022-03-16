package main

import (
	"os/exec"
	"pass"
)

func main() {
	p := pass.Gen(16)

	cmd := exec.Command("xsel", "--input", "--clipboard")
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if _, err := in.Write([]byte(p)); err != nil {
		panic(err)
	}
	if err := in.Close(); err != nil {
		panic(err)
	}
	cmd.Wait()
}
