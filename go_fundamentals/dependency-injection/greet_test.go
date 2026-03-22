package dependency_injection_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	. "github.com/KaelDucatti/go-with-tests/go_fundamentals/dependency-injection"
)

func TestGreetSuccessCases(t *testing.T) {
	t.Run("should return hello to buffer", func(t *testing.T) {
		require := require.New(t)
		buffer := bytes.Buffer{}
		Greet(&buffer, "Kael")

		expected := "Hello, Kael"
		actual := buffer.String()

		require.Equal(expected, actual)
	})
}
