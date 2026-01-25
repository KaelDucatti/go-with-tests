package structs_methods_intefaces

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Shape interface {
	Area() 		float32
	Perimeter() float32
}

type Rectangle struct {
	Height 	float32
	Width	float32
}

type Square struct {
	Edge 	float32
}

type Circle struct {
	Radius	float32
}

func CheckPerimeter(t *testing.T, shape Shape, expected float32) {
	t.Helper()
	require := require.New(t)

	actual := shape.Area()

	require.Equal(expected, actual)
}