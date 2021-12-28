package main

import (
	"fmt"
	"secret"
)

func main() {
	v := secret.New("my-fake-key")

	if err := v.Set("demo_key", "some crazy value"); err != nil {
		panic(err)
	}

	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain: ", plain)
}
