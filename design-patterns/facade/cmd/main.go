package main

import (
	"facade/email"
	"facade/storage"
	"facade/validate"
	"fmt"
	"log"
)

func main() {
	showInit()

	tkn := "valid"
	user := "blogger"
	to := "J@Algo.net"
	cmmt := "Well done!"

	vt := validate.NewValidatorToken(tkn)
	if err := vt.Validate(); err != nil {
		log.Fatal(err)
	}

	vp := validate.NewValidatorPermission(user)
	if err := vp.Validate(); err != nil {
		log.Fatal(err)
	}

	s := storage.New("postgres")
	s.Save(cmmt)

	e := email.New()
	e.Send(to, cmmt)

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
