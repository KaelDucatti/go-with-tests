package structs_methods_intefaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRectangle(t *testing.T) {
	t.Run("Success Cases", func (t *testing.T) {
		t.Run("should return the rectangle perimeter", func (t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(5, 4)
			
			expected := float32(18)
			actual := r.Perimeter()
	
			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should return the rectangle area", func (t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(5, 4)
	
			expected := float32(20)
			actual := r.Area()
	
			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
	t.Run("Valition Errors", func (t *testing.T) {
		t.Run("should return error when values are zero or negative", func (t *testing.T) {
			require := require.New(t)
			
			testCases := []struct {
				name	string
				height	float32
				width 	float32
			}{
				{"height is zero", 0, 4},
				{"width is zero", 5, 0},
				{"height is negative", -5, 4},
				{"width is negative", 5, -4},
			}
	
			for _, tc := range testCases {
				t.Run(tc.name, func (t *testing.T) {
					_, err := NewRectangle(tc.height, tc.width)
					require.Error(err)
					require.EqualError(err, "neither value must be equal to or less than 0")
				})
			}
		})
	})
}

func ExampleNewRectangle() {
	r, _ := NewRectangle(5, 4)
	fmt.Println(r)
	// Output: &{5 4}
}

func ExampleRectangle_Perimeter() {
	r, _ := NewRectangle(5, 4)
	fmt.Println(r.Perimeter())
	// Output: 18
}

func ExampleRectangle_Area() {
	r, _ := NewRectangle(5, 4)
	fmt.Println(r.Area())
	// Output: 20
}

func BenchmarkRectangle(b *testing.B) {
	for b.Loop() {
		NewRectangle(5, 4)
	}
}

func BenchmarkRectangle_Perimeter(b *testing.B) {
	r, _ := NewRectangle(5, 4)
	for b.Loop() {
		r.Perimeter()
	}
}

func BenchmarkRectangle_Area(b *testing.B) {
	r, _ := NewRectangle(5, 4)
	for b.Loop() {
		r.Area()
	}
}