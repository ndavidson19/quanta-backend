package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTx(t *testing.T) {
	store := NewStore(testDB)

	account := CreateAccountParams{
		Owner:    "test",
		Balance:  0,
		Currency: "USD",
	}

	// run n concurrent transactions
	n := int64(2)
	ammount := int64(10)

	errs := make(chan error)
	results := make(chan CreateTxResult)

	fmt.Println(">> before: ", account.Balance)

	for i := int64(0); i < n; i++ {
		txName := fmt.Sprintf("tx %d", i+1)

		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			result, err := store.CreateTx(context.Background(), CreateTxParams{
				Owner:  account.Owner,
				Amount: ammount,
			})

			errs <- err
			results <- result
			ctx.Done()
		}()
	}

	existed := make(map[int64]bool)
	// check results
	for i := int64(0); i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check account
		require.NotEmpty(t, result.Account)
		require.Equal(t, account.Owner, result.Account.Owner)
		require.Equal(t, account.Balance+ammount, result.Account.Balance)
		require.Equal(t, account.Currency, result.Account.Currency)

		// check deposit
		require.NotEmpty(t, result.Deposit)
		require.Equal(t, account.ID, result.Deposit.AccountID)
		require.Equal(t, ammount, result.Deposit.Amount)

		// check audit log
		require.NotEmpty(t, result.AuditLog)
		require.Equal(t, account.ID, result.AuditLog.AccountID)
		require.Equal(t, "deposit", result.AuditLog.Action)
		require.NotEmpty(t, result.AuditLog.Timestamp)

		// check balance
		fmt.Println(">> balance: ", result.Balance)
		require.NotEmpty(t, result.Balance)
		require.Equal(t, account.Balance+ammount, result.Balance)
		require.True(t, result.Balance >= 0)
		require.True(t, result.Balance%ammount == 0)

		k := int64(result.Balance / ammount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true

		_, err = store.GetAccount(context.Background(), result.Account.ID)
		require.NoError(t, err)

	}

	// check final account
	updatedAccount, err := store.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.Balance+n*ammount, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	fmt.Println(">> after: ", updatedAccount.Balance)

}
