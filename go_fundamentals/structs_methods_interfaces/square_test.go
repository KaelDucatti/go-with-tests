package structs_methods_intefaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)



func TestSquare(t *testing.T) {
	t.Run("Success Cases", func (t *testing.T) {
		t.Run("should return the square perimeter", func (t *testing.T) {
			require := require.New(t)
			s, err := NewSquare(5)

			expected := float32(20)
			actual := s.Perimeter()

			require.NoError(err)
			require.Equal(expected, actual)
			
		})
		t.Run("should return the square area", func (t *testing.T) {
			require := require.New(t)
			s, err := NewSquare(5)

			expected := float32(25)
			actual := s.Area()

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

func ExampleNewSquare() {
	s, _ := NewSquare(5)
	fmt.Println(s)
	// Output: &{5}
}

func ExampleSquare_Perimeter() {
	s, _ := NewSquare(5)
	fmt.Println(s.Perimeter())
	// Output: 20
}

func ExampleSquare_Area() {
	s, _ := NewSquare(5)
	fmt.Println(s.Area())
	// Output: 25
}

func BenchmarkSquare(b *testing.B) {
	for b.Loop() {
		NewSquare(5)
	}
}

func BenchmarkSquare_Perimeter(b *testing.B) {
	s, _ := NewSquare(5)
	for b.Loop() {
		s.Perimeter()
	}
}

func BenchmarkSquare_Area(b *testing.B) {
	s, _ := NewSquare(5)
	for b.Loop() {
		s.Area()
	}
}