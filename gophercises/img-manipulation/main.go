package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
  out, err := primitive("tmp/jump.png", "out.png", 100, triangle)
	cmd := exec.Command("primitive", "args")
	b, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

