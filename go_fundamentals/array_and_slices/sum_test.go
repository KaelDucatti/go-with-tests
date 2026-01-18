package array_and_slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("should return the array summarize", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		slice := []int{1, 2, 3, 4, 5}

		exepected := 10
		actual, err := Sum(slice)

		require.NoError(err)
		assert.Equal(exepected, actual)
	})
}