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
	})
}
