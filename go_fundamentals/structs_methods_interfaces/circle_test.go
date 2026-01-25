package structs_methods_intefaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircle(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return the circle perimeter", func (t *testing.T) {
			require := require.New(t)
			c, err := NewCircle(10)

			expected := float32(62.831856)
			actual := c.Perimeter()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should return the circle area", func (t *testing.T) {
			require := require.New(t)
			c, err := NewCircle(10)

			expected := float32(314.15927)
			actual := c.Area()

			require.NoError(err)
			require.Equal(expected, actual)
		})
	})
	t.Run("Validation Errors", func (t *testing.T) {
		require := require.New(t)

		testCases := []struct {
			name	string
			radius	float32
		}{
			{"radius is zero", 0},
			{"radius is negative", -10},
		}

		for _, test := range testCases {
			t.Run(test.name, func (t *testing.T) {
				_, err := NewCircle(test.radius)
				require.Error(err)
				require.EqualError(err, "radius value cannot be equal to or less than 0")
			})
		}
	})
}

func ExampleNewCircle() {
	c, _ := NewCircle(10)
	fmt.Println(c)
	// Output: &{10}
}

func ExampleCircle_Perimeter() {
	c, _ := NewCircle(10)
	fmt.Println(c.Perimeter())
	// Output: 62.831856
}

func ExampleCircle_Area() {
	c, _ := NewCircle(10)
	fmt.Println(c.Area())
	// Output: 314.15927
}

func BenchmarkCircle(b *testing.B) {
	for b.Loop() {
		NewCircle(10)
	}
}

func BenchmarkCircle_Perimeter(b *testing.B) {
	c, _ := NewCircle(10)
	for b.Loop() {
		c.Perimeter()
	}
}

func BenchmarkCircle_Area(b *testing.B) {
	c, _ := NewCircle(10)
	for b.Loop() {
		c.Area()
	}
}