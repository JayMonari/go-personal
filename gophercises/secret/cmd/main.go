package main

import (
	"fmt"
	"secret"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")

	if err := v.Set("demo_key", "some crazy value"); err != nil {
		panic(err)
	}

	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain: ", plain)
}
