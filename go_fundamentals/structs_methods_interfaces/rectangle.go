package structs_methods_intefaces

import (
	"errors"
)

type Rectangle struct {
	Height 	float32
	Width	float32
}

func NewRectangle(height, width float32) (*Rectangle, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("neither value must be equal to or less than 0")
	}

	return &Rectangle{Height: height, Width: width}, nil
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

func (r *Rectangle) Area() float32 {
	return r.Height * r.Width
}