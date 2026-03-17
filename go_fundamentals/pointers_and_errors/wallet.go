package pointers_and_errors

const (
	ErrInputIsNotPositive = ErrWallet("Input need to be a positeve number")
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

func (w *Wallet) Deposit(balance float32) error {
	if err := ValidateIfBalanceIsPositive(balance); err != nil {
		return err
	}
	w.balance = Bitcoin(balance)
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func ValidateIfBalanceIsPositive(balance float32) error {
	if balance <= 0 {
		return ErrInputIsNotPositive
	}
	return nil
}
