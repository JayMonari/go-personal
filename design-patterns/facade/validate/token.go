package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("the user is not logged in")

type Token struct{ token string }

func NewToken(t string) Token {
	return Token{token: t}
}

func (vt *Token) Validate() error {
	if vt.token != "valid" {
		return ErrTokenNotValid
	}

	fmt.Println("valid token")
	return nil
}
