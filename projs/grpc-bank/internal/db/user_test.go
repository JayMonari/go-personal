package db_test

import (
	"context"
	"testing"
	"time"

	"example.xyz/bank/internal/db"
	"example.xyz/bank/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) db.User {
	hashedPass, err := util.HashPassword(util.RandString(6))
	require.NoError(t, err)
	arg := db.CreateUserParams{
		Username:       util.RandOwner(),
		HashedPassword: hashedPass,
		FullName:       util.RandOwner(),
		Email:          util.RandEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotEmpty(t, user)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.NotEmpty(t, user2)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, 9*time.Millisecond)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, 9*time.Millisecond)
}

// func TestUpdateAccount(t *testing.T) {
// 	acc1 := createRandomAccount(t)
// 	arg := db.UpdateAccountParams{
// 		ID:      acc1.ID,
// 		Balance: util.RandBalance(),
// 	}
// 	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
//
// 	require.NotEmpty(t, acc2)
// 	require.Equal(t, arg.Balance, acc2.Balance)
// 	require.Equal(t, acc1.Currency, acc2.Currency)
// 	require.Equal(t, acc1.ID, acc2.ID)
// 	require.Equal(t, acc1.Owner, acc2.Owner)
// 	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, 9*time.Millisecond)
// }
//
// func TestDeleteAccount(t *testing.T) {
// 	acc1 := createRandomAccount(t)
// 	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
// 	require.NoError(t, err)
//
// 	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, acc2)
// }
//
// func TestListAccounts(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomAccount(t)
// 	}
//
// 	arg := db.ListAccountsParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}
// 	accounts, err := testQueries.ListAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.Len(t, accounts, 5)
//
// 	for _, a := range accounts {
// 		require.NotEmpty(t, a)
// 	}
// }
