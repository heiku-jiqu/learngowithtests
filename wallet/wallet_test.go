package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
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
		err := wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, e error, want error) {
	t.Helper()
	if e == nil {
		t.Fatal("didnt get an error") // stops test execution
	}
	if e != want {
		t.Errorf("got %q, want %q", e, want)
	}
}

func assertNoError(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Fatal("expected no error")
	}
}
