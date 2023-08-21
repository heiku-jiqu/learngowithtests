package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amt int) {
	w.balance += amt
}

func (w *Wallet) Balance() int {
	return w.balance
}

// arguments of functions are COPIED (passed by value)
func (w Wallet) BrokenDeposit(amt int) {
	fmt.Printf("address of w.balance in Deposit() is %v \n", &w.balance)
	w.balance += amt
}
