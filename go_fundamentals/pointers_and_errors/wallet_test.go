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
			err := wallet.Deposit(Bitcoin(10))

			expected := Bitcoin(10)
			actual := wallet.Balance()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should return balance equals to 5", func(t *testing.T) {
			require := require.New(t)
			wallet := &Wallet{balance: 10}
			err := wallet.Withdraw(Bitcoin(5))

			expected := Bitcoin(5)
			actual := wallet.Balance()

			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
	t.Run("Validation Error", func(t *testing.T) {
		t.Run("should fail when deposit is equals to or less than 0", func(t *testing.T) {
			require := require.New(t)
			wallet := NewWallet()

			testCases := []struct {
				testName string
				balance  Bitcoin
			}{
				{"Bitcoin is 0", Bitcoin(0)},
				{"Bitcoin is negative", Bitcoin(-1)},
			}

			for _, test := range testCases {
				t.Run(test.testName, func(t *testing.T) {
					err := wallet.Deposit(test.balance)
					require.Error(err)
					require.EqualError(err, "deposit value must be equals to or greater than 0")
				})
			}
		})
		t.Run("should fail when withdraw is greater than balance", func(t *testing.T) {
			require := require.New(t)
			wallet := &Wallet{balance: Bitcoin(10)}

			err := wallet.Withdraw(Bitcoin(11))

			require.Error(err)
			require.EqualError(
				err, "the withdraw value connot be greater than the actual balance",
			)
		})
		t.Run("should fail when withdraw is equals to or less than 0", func(t *testing.T) {
			require := require.New(t)
			wallet := &Wallet{balance: Bitcoin(10)}

			testCases := []struct {
				testName string
				balance  Bitcoin
			}{
				{"Withdraw is 0", 0},
				{"Withdraw is negative", -1},
			}

			for _, test := range testCases {
				err := wallet.Withdraw(test.balance)
				require.Error(err)
				require.EqualError(err, "the withdraw value cannot be equals to or less than 0")
			}
		})
	})
}

func ExampleWallet_Deposit() {
	wallet := &Wallet{balance: 0}
	_ = wallet.Deposit(Bitcoin(10))
	fmt.Println(wallet.Balance())
	// Output: 10
}

func ExampleWallet_Withdraw() {
	wallet := &Wallet{balance: Bitcoin(10)}
	_ = wallet.Withdraw(Bitcoin(5))
	fmt.Println(wallet.Balance())
	// Output: 5
}

func BenchmarkWallet_Deposit(b *testing.B) {
	for b.Loop() {
		wallet := &Wallet{}
		_ = wallet.Deposit(Bitcoin(1))
	}
}

func BenchmarkWallet_Withdraw(b *testing.B) {
	for b.Loop() {
		wallet := &Wallet{balance: Bitcoin(1000)}
		_ = wallet.Withdraw(Bitcoin(1))
	}
}
