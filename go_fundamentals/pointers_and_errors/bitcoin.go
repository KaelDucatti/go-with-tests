package pointers_and_errors

import "fmt"

type Bitcoin float32

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}