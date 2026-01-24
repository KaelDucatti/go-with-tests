package structs_methods_intefaces

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCircle(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return the circle perimeter", func (t *testing.T) {
			require := require.New(t)
			c, err := NewCircle(10)

			expected := float32(62.83)
			actual := c.Perimeter()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should return the circle area", func (t *testing.T) {
			require := require.New(t)
			c, err := NewCircle(10)

			expected := float32(314.15)
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