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
		require := require.New(t)
		slice := []int{1, 2, 3, 4, 5}

		exepected := 15
		actual, err := Sum(slice)

		require.NoError(err)
		assert.Equal(exepected, actual)
	})
	t.Run("list must have at least 1 value", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		slice := []int{}

		_, err := Sum(slice)

		require.Error(err)
		assert.EqualError(err, "the list must have at least 1 value")
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should return an array with the total of N slices", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

		expected := []int{6, 15, 24}
		actual, err := SumAll(matrix)

		require.NoError(err)
		assert.Equal(expected, actual)
	})
	t.Run("all lists must hava at least 1 value", func (t *testing.T){
		assert := assert.New(t)
		require := require.New(t)
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

		_, err := SumAll(matrix)

		require.Error(err)
		assert.EqualError(err, "all lists must have at least 1 value")
	})
}

func ExampleSum() {
	total, _ := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(total)
	// Output: 15
}

func ExampleSumAll() {
	total, _ := SumAll([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(total)
	// Output: []int{6, 15, 24}
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}