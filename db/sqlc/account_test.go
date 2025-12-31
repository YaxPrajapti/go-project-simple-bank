package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yax/simple-bank/util"
)

func createRandomAccount(t *testing.T) Account {
	createAccountParams := CreateAccountParams{
		AccountHolderName: util.RandomOwner(),
		Balance:           util.RandomMoney(),
		Currency:          util.RandomCurrency(),
	}
	acc, err := testQueris.CreateAccount(context.Background(), createAccountParams)
	require.NoError(t, err, "CreateAccount should not fail")
	require.NotEmpty(t, acc)

	require.Equal(t, createAccountParams.AccountHolderName, acc.AccountHolderName)
	require.Equal(t, createAccountParams.Balance, acc.Balance)
	require.Equal(t, createAccountParams.Currency, acc.Currency)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// create account
	acc1 := createRandomAccount(t)
	acc2, err := testQueris.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc2.ID, acc1.ID)
	require.Equal(t, acc2.AccountHolderName, acc1.AccountHolderName)
	require.Equal(t, acc2.Balance, acc1.Balance)
	require.Equal(t, acc2.Currency, acc1.Currency)
}

func TestUpdateAccountBalance(t *testing.T) {
	acc1 := createRandomAccount(t)
	arg := UpdateAccountBalanceParams{
		ID:      acc1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueris.UpdateAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc2.ID, acc1.ID)
	require.Equal(t, acc2.AccountHolderName, acc1.AccountHolderName)
	require.Equal(t, acc2.Balance, arg.Balance)
	require.Equal(t, acc2.Currency, acc1.Currency)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	err := testQueris.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := testQueris.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error()[5:])
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	for range 10 {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueris.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)

	}
}
