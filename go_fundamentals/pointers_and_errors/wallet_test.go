package pointers_and_errors_test

import (
	"testing"

	. "github.com/KaelDucatti/go-with-tests/go_fundamentals/pointers_and_errors"
	"github.com/stretchr/testify/require"
)

func TestWallet(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return an wallet instance", func(t *testing.T) {
			require := require.New(t)
			w := NewWallet()

			expected := Bitcoin(0)
			actual := w.Balance()	

			require.Equal(expected, actual)
		})
		t.Run("should deposit 10 bitcoins in the wallet", func(t *testing.T) {
			require := require.New(t)
			w := NewWallet()
			err := w.Deposit(10)

			expected := Bitcoin(10)
			actual := w.Balance()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should withdraw 10 bitcoins from thw wallet", func(t *testing.T) {
			require := require.New(t)
			w := NewWallet()
			w.Deposit(10)

			err := w.Withdraw(5)
			expected := Bitcoin(5)
			actual := w.Balance()

			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
}
