package countdown_test

import (
	"bytes"
	"testing"

	. "github.com/KaelDucatti/go-with-tests/go_fundamentals/mocking"
	"github.com/stretchr/testify/require"
)

func TestCountdown(t *testing.T) {
	require := require.New(t)
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	expected := "1\n2\n3\nGo\n"
	actual := buffer.String()

	require.Equal(expected, actual)
}
