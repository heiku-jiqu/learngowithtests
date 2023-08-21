package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.BrokenDeposit(Bitcoin(10)) // doesn't do anything since passed by val!
		fmt.Printf("address of wallet.balance in test is %v \n", &wallet.balance)

		wallet.Deposit(10)

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
}
