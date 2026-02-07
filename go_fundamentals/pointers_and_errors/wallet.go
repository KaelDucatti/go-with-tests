package pointers_and_errors

import "errors"

type Wallet struct {
	balance Bitcoin
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) Deposit(value Bitcoin) error {
	if value <= 0 {
		return errors.New("deposit value must be equals to or greater than 0")
	}
	w.balance += value
	return nil
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return errors.New("the withdraw value connot be greater than the actual balance")
	}
	if value <= 0 {
		return errors.New("the withdraw value cannot be equals to or less than 0")
	}
	w.balance -= value
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
