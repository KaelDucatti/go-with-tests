package countdown_test

import (
	"bytes"
	"testing"

	. "github.com/KaelDucatti/go-with-tests/go_fundamentals/mocking"
	"github.com/stretchr/testify/require"
)

func TestCountdown(t *testing.T) {
	t.Run("shold return the countdown", func(t *testing.T) {
		require := require.New(t)
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}
	
		Countdown(buffer, sleeper)
		expected := "3\n2\n1\nGo\n"
		actual := buffer.String()
	
		require.Equal(expected, actual)
		require.Equal(3, sleeper.Calls)
	})

	t.Run("should return the SpyCountdownOperation list", func(t *testing.T) {
		require := require.New(t)
		sco := &SpyCountdownOperations{}
		Countdown(sco, sco)

		expected := []string {
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
		}
		actual := sco.Calls

		require.Equal(expected, actual)
	})
}
