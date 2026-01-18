package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	t.Run("should print 'a' 5 times", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		expected := "aaaaa"
		actual, err := Repeat("a", 5)

		require.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("should return 'value cannot be empty'", func (t *testing.T){
		require := require.New(t)

		expected := "value cannot be empty"
		_, err := Repeat("", 5)

		require.Error(err)
		require.EqualError(err, expected)
	})
	t.Run("shoud return 'count must be greater than 0'", func (t *testing.T){
		require := require.New(t)

		expected := "count must be greater than 0"
		_, err := Repeat("a", 0)

		require.Error(err)
		require.EqualError(err, expected)
	})
}

func ExampleRepeat() {
	value, _ := Repeat("a", 5)
	fmt.Println(value)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}