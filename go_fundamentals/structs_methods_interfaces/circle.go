package structs_methods_intefaces

import (
	"errors"
	"math"
)


func NewCircle(radius float32) (*Circle, error) {
	if radius <= 0 {
		return nil, errors.New("radius value cannot be equal to or less than 0")
	}

	return &Circle{Radius: radius}, nil
}

func (c *Circle) Perimeter() float32 {
	return math.Pi * (2 * c.Radius)
}

func (c *Circle) Area() float32 {
	return math.Pi * (c.Radius * c.Radius)
}