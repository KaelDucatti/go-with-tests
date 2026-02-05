package pointers_and_errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return balance equals to 10", func(t *testing.T) {
			require := require.New(t)
			wallet := NewWallet()
			err := wallet.Deposit(10)

			expected := float64(10)
			actual := wallet.Balance()

			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
}

func ExampleWallet_Deposit() {
	wallet := NewWallet()
	_ = wallet.Deposit(10)
	fmt.Println(wallet.Balance())
	// Output: 10
}

func BenchmarkWallet_Deposit(b *testing.B) {
	wallet := NewWallet()
	for b.Loop() {
		_ = wallet.Deposit(10)
	}
}
