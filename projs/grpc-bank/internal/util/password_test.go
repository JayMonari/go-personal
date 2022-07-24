package util_test

import (
	"testing"

	"example.xyz/bank/internal/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	p := util.RandString(6)
	hashed, err := util.HashPassword(p)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)
	require.NoError(t, util.CheckPassword(p, hashed))

	require.EqualError(t, // incorrect pass
		util.CheckPassword(util.RandString(10), hashed),
		bcrypt.ErrMismatchedHashAndPassword.Error())
}
