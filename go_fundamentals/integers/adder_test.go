package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
) 

func TestAdder(t *testing.T) {
	t.Run("Success Cases", func (t *testing.T){
		t.Run("should return (int, nil)", func (t *testing.T){
			assert := assert.New(t)
			require := require.New(t)
			
			sum, err := Adder(2, 2)
			expected := 4
	
			require.NoError(err)		
			assert.Equal(expected, sum)
		}) 
	})

	t.Run("Validation Errors", func (t *testing.T){
		t.Run("should return 'Num1 is required (connot be zero)'", func (t *testing.T){
			assert := assert.New(t)
			require := require.New(t)
	
			_, err := Adder(0, 2)
			
			
			require.Error(err)
			assert.EqualError(err, "Num1 is required (cannot be zero).")
		})
		t.Run("should return 'Num2 is required (cannot be zero)'", func (t *testing.T){
			assert := assert.New(t)
			require := require.New(t)

			_, err := Adder(2, 0)

			require.Error(err)
			assert.EqualError(err, "Num2 is required (cannot be zero).")
		})
	})
}

func ExampleAdder() {
	sum, _ := Adder(2, 2)
	fmt.Println(sum)
	// Output: 4
}

func BenchmarkAdder(b *testing.B) {
	for b.Loop() {
		Adder(2, 2)
	}
}