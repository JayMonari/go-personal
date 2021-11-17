package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("the user is not logged in")

type ValidatorToken struct{ token string }

func NewValidatorToken(t string) ValidatorToken {
	return ValidatorToken{token: t}
}

func (vt *ValidatorToken) Validate() error {
	if vt.token != "valid" {
		return ErrTokenNotValid
	}

	fmt.Println("valid token")
	return nil
}
