package db_test

import (
	"context"
	"testing"

	"example.xyz/bank/internal/db"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := db.NewStore(testDB)

	acct1 := createRandomAccount(t)
	acct2 := createRandomAccount(t)
	t.Log(">> before:", acct1.Balance, acct2.Balance)

	n := 5
	amount := int64(10)
	errs := make(chan error)
	results := make(chan db.TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			res, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: acct1.ID,
				ToAccountID:   acct2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- res
		}()
	}

	seen := map[int]struct{}{}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		res := <-results
		require.NotEmpty(t, res)

		// check transfer
		trans := res.Transfer
		require.NotEmpty(t, trans)
		require.Equal(t, acct1.ID, trans.FromAccountID)
		require.Equal(t, acct2.ID, trans.ToAccountID)
		require.Equal(t, amount, trans.Amount)
		require.NotZero(t, trans.ID)
		require.NotZero(t, trans.CreatedAt)

		_, err = store.GetTransfer(context.Background(), trans.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := res.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, acct1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := res.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, acct2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := res.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, acct1.ID, fromAccount.ID)

		toAccount := res.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, acct2.ID, toAccount.ID)

		// check accounts' balance
		t.Log(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := acct1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - acct2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)
		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, seen, k)
		seen[k] = struct{}{}
	}

	// check final updated balance
	updatedAcct1, err := testQueries.GetAccount(context.Background(), acct1.ID)
	require.NoError(t, err)
	updatedAcct2, err := testQueries.GetAccount(context.Background(), acct2.ID)
	require.NoError(t, err)

	require.Equal(t, acct1.Balance-int64(n)*amount, updatedAcct1.Balance)
	require.Equal(t, acct2.Balance+int64(n)*amount, updatedAcct2.Balance)
	t.Log(">> after:", updatedAcct1.Balance, updatedAcct2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := db.NewStore(testDB)

	acct1 := createRandomAccount(t)
	acct2 := createRandomAccount(t)
	t.Log(">> before:", acct1.Balance, acct2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID, toAccountID := acct1.ID, acct2.ID
		if i%2 == 1 {
			fromAccountID = acct2.ID
			toAccountID = acct1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}
	// check final updated balance
	updatedAcct1, err := testQueries.GetAccount(context.Background(), acct1.ID)
	require.NoError(t, err)
	updatedAcct2, err := testQueries.GetAccount(context.Background(), acct2.ID)
	require.NoError(t, err)

	require.Equal(t, acct1.Balance, updatedAcct1.Balance)
	require.Equal(t, acct2.Balance, updatedAcct2.Balance)
	t.Log(">> after:", updatedAcct1.Balance, updatedAcct2.Balance)
}
