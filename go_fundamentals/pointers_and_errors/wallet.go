package pointers_and_errors

import "errors"

type Wallet struct {
	balance float64
}

func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) Deposit(value float64) error {
	if value <= 0 {
		return errors.New("deposit value must be equals to or greater than 0")
	}
	w.balance += value
	return nil
}

func (w *Wallet) Balance() float64 {
	return w.balance
}
