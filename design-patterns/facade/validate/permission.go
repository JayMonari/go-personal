package validate

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("the user is not authenticated to comment on this post")

type Permission struct{ userID string }

func NewPermission(ID string) Permission {
	return Permission{userID: ID}
}

func (vp *Permission) Validate() error {
	if vp.userID != "blogger" {
		return ErrPermissionNotValid
	}

	fmt.Println("user is authenticated to comment")
	return nil
}
