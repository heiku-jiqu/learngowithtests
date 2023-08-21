package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.BrokenDeposit(10) // doesn't do anything since passed by val!
	fmt.Printf("address of wallet.balance in test is %v \n", &wallet.balance)

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
