package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

type SignUpPayload struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,containsany=!@#?*"`
	Name     string `validate:"omitempty,min=4"`
}

func main() {
	v := validator.New()
	if err := v.Struct(SignUpPayload{
		Email:    "test@testcom",
		Password: "suhcrpas",
		Name:     "123",
	}); err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Fatal(err)
		}
		for _, verr := range verrs {
			fmt.Printf("%q has a value of '%v' which does not satisfy %q.\n", verr.Field(), verr.Value(), verr.Tag())
		}
	}

	if err := v.Var("gewdpassword", "required,min=8,containsany=!@#$?*"); err != nil {
		fmt.Println(err)
	}
}
