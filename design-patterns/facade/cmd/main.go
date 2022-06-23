package main

import (
	"facade/facade"
	"fmt"
	"log"
)

func main() {
	showInit()

	tkn := "valid"
	user := "blogger"
	to := "J@Algo.net"
	cmmt := "Well done!"
	fac := facade.New(to, cmmt, tkn, user)
	err := fac.Comment()
	if err != nil {
		log.Fatal(err)
	}

  fac.Notify()

	showFinish()
}

func showInit() {
	fmt.Print(`
***********************
* Welcome to the Blog *
***********************
`)
}

func showFinish() {
	fmt.Print(`
**************************
* Thanks for Commenting! *
**************************
`)
}
