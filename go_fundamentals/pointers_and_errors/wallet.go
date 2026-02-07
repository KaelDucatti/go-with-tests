package pointers_and_errors

import "errors"

var (
	ErrImpossibleDepositValue error = errors.New(
		"deposit value must be equals to or greater than 0",
	)
	ErrWithdrawGreaterThanBalance error = errors.New(
		"the withdraw value connot be greater than the actual balance",
	)
	ErrWithdrawEqualLessThanZero error = errors.New(
		"the withdraw value cannot be equals to or less than 0",
	)
)

type Wallet struct {
	balance Bitcoin
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) Deposit(value Bitcoin) error {
	if value <= 0 {
		return ErrImpossibleDepositValue
	}
	w.balance += value
	return nil
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrWithdrawGreaterThanBalance
	}
	if value <= 0 {
		return ErrWithdrawEqualLessThanZero
	}
	w.balance -= value
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
