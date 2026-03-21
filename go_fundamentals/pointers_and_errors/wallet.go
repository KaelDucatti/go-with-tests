package pointers_and_errors

const (
	ErrInputIsNotPositive           = ErrWallet("Input need to be a positeve number")
	ErrWithdrawIsGreaterThanBalance = ErrWallet("Withdraw cannot be greater than the actual balance")
)

type ErrWallet string

func (ew ErrWallet) Error() string {
	return string(ew)
}

type Wallet struct {
	balance Bitcoin
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) Deposit(amount float32) error {
	if err := ValidateIfAmountIsPositive(amount); err != nil {
		return err
	}
	w.balance = Bitcoin(amount)
	return nil
}

func (w *Wallet) Withdraw(amount float32) error {
	if err := ValidateIfAmountIsPositive(amount); err != nil {
		return err
	}
	if err := ValidateIfAmountIsGreaterThanBalance(amount, float32(w.Balance())); err != nil {
		return err
	}
	w.balance -= Bitcoin(amount)
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func ValidateIfAmountIsPositive(amount float32) error {
	if amount <= 0 {
		return ErrInputIsNotPositive
	}
	return nil
}

func ValidateIfAmountIsGreaterThanBalance(amount, balance float32) error {
	if amount > balance {
		return ErrWithdrawIsGreaterThanBalance
	}
	return nil
}