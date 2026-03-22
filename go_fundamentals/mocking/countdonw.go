package countdown

import (
	"fmt"
	"io"
)


func Countdown(w io.Writer) {
	for i := 1; i <= 3; i++ {
		fmt.Fprintln(w, i)
	}
	fmt.Fprintln(w, "Go")
}