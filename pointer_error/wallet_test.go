package wallet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		assert.Equal(t, expected, got)
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expected := Bitcoin(10)

		fmt.Printf("got: %s\n", expected)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)
		assertBalance(t, wallet, expected)
	})
}
