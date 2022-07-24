package db_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"example.xyz/bank/internal/db"
	"example.xyz/bank/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner:    createRandomUser(t).Username,
		Balance:  util.RandBalance(),
		Currency: util.RandCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)
	require.NotEmpty(t, account)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	require.Equal(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.NotEmpty(t, acc2)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, 9*time.Millisecond)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	arg := db.UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandBalance(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, acc2)
	require.Equal(t, arg.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, 9*time.Millisecond)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount db.Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, a := range accounts {
		require.NotEmpty(t, a)
		require.Equal(t, lastAccount.Owner, a.Owner)
	}
}
