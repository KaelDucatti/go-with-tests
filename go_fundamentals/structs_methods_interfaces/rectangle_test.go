package structs_methods_intefaces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRectangle(t *testing.T) {
	t.Run("Success Cases", func(t *testing.T) {
		t.Run("should return rectangle area", func(t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(10, 8)

			expected := float32(80)
			actual := r.Area()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should return rectagle perimeter", func(t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(10, 8)

			expected := float32(36)
			actual := r.Perimeter()

			require.NoError(err)
			require.Equal(expected, actual)
		})
		t.Run("should update the height and width values", func(t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(10, 8)
			require.NoError(r.SetHeight(20))
			require.NoError(r.SetWidth(16))

			expectedHeight := float32(20)
			expectedWidth := float32(16)
			actualHeight := r.Height()
			actualWidth := r.Width()

			require.NoError(err)
			require.Equal(expectedHeight, actualHeight)
			require.Equal(expectedWidth, actualWidth)
		})
		t.Run("Should return the height and width values", func(t *testing.T) {
			require := require.New(t)
			r, err := NewRectangle(10, 8)

			expectedHeight := float32(10)
			expectedWidth := float32(8)
			actualHeight := r.Height()
			actualWidth := r.Width()

			require.NoError(err)
			require.Equal(expectedHeight, actualHeight)
			require.Equal(expectedWidth, actualWidth)
		})
	})
	t.Run("Error Validation", func(t *testing.T) {
		t.Run("should return error when height or width are zero", func(t *testing.T) {
			testCases := []struct {
				title  string
				height float32
				width  float32
			}{
				{"height is zero", 0, 8},
				{"width is zero", 10, 0},
				{"height and width are zero", 0, 0},
			}

			for _, test := range testCases {
				t.Run(test.title, func(t *testing.T) {
					require := require.New(t)
					_, err := NewRectangle(test.height, test.width)

					require.Error(err)
					require.ErrorIs(err, ErrValueIsZero)
				})
			}
		})
		t.Run("should return error when height or width are negative", func(t *testing.T) {
			newRectangleCases := []struct {
				title  string
				height float32
				width  float32
			}{
				{"height is negative", -1, 8},
				{"width is negative", 10, -1},
				{"height and width are negative", -1, -1},
			}

			for _, test := range newRectangleCases {
				t.Run(test.title, func(t *testing.T) {
					require := require.New(t)
					_, err := NewRectangle(test.height, test.width)

					require.Error(err)
					require.ErrorIs(err, ErrValueIsNegative)
				})
			}
		})
		t.Run("should return error when update value are negative", func(t *testing.T) {
			require := require.New(t)

			r, _ := NewRectangle(10, 8)

			errHeight := r.SetHeight(-10)
			errWidth := r.SetWidth(-8)

			require.ErrorIs(errHeight, ErrValueIsNegative)
			require.ErrorIs(errWidth, ErrValueIsNegative)
		})
		t.Run("should return error when update value are zero", func(t *testing.T) {
			require := require.New(t)

			r, _ := NewRectangle(10, 8)

			errHeight := r.SetHeight(0)
			errWidth := r.SetWidth(0)

			require.ErrorIs(errHeight, ErrValueIsZero)
			require.ErrorIs(errWidth, ErrValueIsZero)
		})
		t.Run("should return error when height and widht are equal", func(t *testing.T) {
			require := require.New(t)

			_, err := NewRectangle(10, 10)

			require.Error(err)
			require.ErrorIs(err, ErrHeightAndWidthAreEquals)
		})
		t.Run("should return error when update value is equal", func(t *testing.T) {
			require := require.New(t)
			r, _ := NewRectangle(10, 8)

			errHeight := r.SetHeight(8)
			errWidth := r.SetWidth(10)

			require.ErrorIs(errHeight, ErrHeightAndWidthAreEquals)
			require.ErrorIs(errWidth, ErrHeightAndWidthAreEquals)
		})
	})
}

func ExampleNewRectangle() {
	s, _ := NewRectangle(10, 8)
	fmt.Println(s)
	// Output: &{10 8}
}

func ExampleRectangle_Area() {
	s, _ := NewRectangle(10, 8)
	fmt.Println(s.Area())
	// Output: 80
}

func ExampleRectangle_Perimeter() {
	s, _ := NewRectangle(10, 8)
	fmt.Println(s.Perimeter())
	// Output: 36
}

func ExampleRectangle_SetHeight() {
	s, _ := NewRectangle(10, 8)
	s.SetHeight(20)
	fmt.Println(s.Height())
	// Output: 20
}

func ExampleRectangle_SetWidth() {
	s, _ := NewRectangle(10, 8)
	s.SetWidth(16)
	fmt.Println(s.Width())
	// Output: 16
}

func BenchmarkRectangle(b *testing.B) {
	for b.Loop() {
		_, _ = NewRectangle(10, 8)
	}
}

func BenchmarkRectangle_Area(b *testing.B) {
	s, _ := NewRectangle(10, 8)
	for b.Loop() {
		_ = s.Area()
	}
}

func BenchmarkRectangle_Perimeter(b *testing.B) {
	s, _ := NewRectangle(10, 8)
	for b.Loop() {
		_ = s.Perimeter()
	}
}

func BenchmarkRectangle_SetHeight(b *testing.B) {
	for	b.Loop() {
		s, _ := NewRectangle(10, 8)
		s.SetHeight(20)
	} 
}

func BenchmarkRectangle_SetWidth(b *testing.B) {
	for	b.Loop() {
		s, _ := NewRectangle(10, 8)
		s.SetWidth(16)
	} 
}