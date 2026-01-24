package structs_methods_intefaces

import (
	"testing"

	"github.com/stretchr/testify/require"
)



func TestSquare(t *testing.T) {
	t.Run("Success Cases", func (t *testing.T) {
		t.Run("should return the square perimeter", func (t *testing.T) {
			require := require.New(t)
			s, err := NewSquare(5)

			expected := float32(20)
			actual := Perimeter(s)

			require.NoError(err)
			require.Equal(expected, actual)
			
		})
		t.Run("should return the square area", func (t *testing.T) {
			require := require.New(t)
			s, err := NewSquare(5)

			expected := float32(25)
			actual := Area(s)

			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
	t.Run("Validation Erros", func (t *testing.T) {
		t.Run(
			"should return error when edge is equal to or less than 0", 
			func (t *testing.T) {
				require := require.New(t)
				
				testCases := []struct{
					name 	string
					edge	float32
				}{
					{"edge is zero", 0},
					{"edge is negative", -5},
				}

				for _, test := range testCases {
					t.Run(test.name, func (t *testing.T) {
						_, err := NewSquare(test.edge)
						require.Error(err)
						require.EqualError(err, "value must be equal to or greater than 0")
					})
				}
		})
	})
}