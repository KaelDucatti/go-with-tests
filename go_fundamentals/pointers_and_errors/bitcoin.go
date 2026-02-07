package pointers_and_errors

import "fmt"

type Stringer interface {
	String() string
}

type Bitcoin float64

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC\n", b)
}
