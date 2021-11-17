package validate

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("the user is not authenticated to comment on this post")

type ValidatorPermission struct{ userID string }

func NewValidatorPermission(ID string) ValidatorPermission {
	return ValidatorPermission{userID: ID}
}

func (vp *ValidatorPermission) Validate() error {
	if vp.userID != "blogger" {
		return ErrPermissionNotValid
	}

	fmt.Println("user is authenticated to comment")
	return nil
}
