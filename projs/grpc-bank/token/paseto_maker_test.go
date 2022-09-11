package token_test

import (
	"testing"
	"time"

	"example.xyz/bank/internal/util"
	"example.xyz/bank/token"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	m, err := token.NewPasetoMaker(util.RandString(32))
	require.NoError(t, err)

	username := util.RandOwner()
	dur := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(dur)

	tok, payload, err := m.CreateToken(username, dur)
	require.NoError(t, err)
	require.NotEmpty(t, tok)
	require.NotEmpty(t, payload)

	payload, err = m.VerifyToken(tok)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	m, err := token.NewPasetoMaker(util.RandString(32))
	require.NoError(t, err)

	tok, payload, err := m.CreateToken(util.RandOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, tok)
	require.NotEmpty(t, payload)

	payload, err = m.VerifyToken(tok)
	require.EqualError(t, err, token.ErrExpiredToken.Error())
	require.Nil(t, payload)
}
