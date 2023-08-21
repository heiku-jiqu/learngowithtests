package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient balance")

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
	}
)

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// arguments of functions are COPIED (passed by value)
func (w Wallet) BrokenDeposit(amt Bitcoin) {
	fmt.Printf("address of w.balance in Deposit() is %v \n", &w.balance)
	w.balance += amt
}
