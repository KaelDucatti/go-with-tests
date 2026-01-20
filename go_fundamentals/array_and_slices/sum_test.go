package array_and_slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	t.Run("should return the array summarize", func (t *testing.T){
		assert := assert.New(t)
		slice := []int{1, 2, 3, 4, 5}

		expected := 15
		actual := Sum(slice)

		assert.Equal(expected, actual)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should return a slice with the total of N slices", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		expected := []int{6, 15, 24}
		actual, err := SumAll(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		)

		require.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("must have at least 1 list", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		_, err := SumAll()

		require.Error(err)
		assert.EqualError(err, "SumAll must received at least 1 list")
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should return a slice with the sum of all list tails", func (t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		expected := []int{5, 11, 17}
		actual, err := SumAllTails(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		)

		require.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("must have at least 1 list", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)

		_, err := SumAllTails()

		require.Error(err)
		assert.EqualError(err, "SumAllTails must received at least 1 list")
	})
}

func ExampleSum() {
	total := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(total)
	// Output: 15
}

func ExampleSumAll() {
	total, _ := SumAll(
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	)
	fmt.Println(total)
	// Output: [6 15 24]
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll(
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},	
		)
	}
}