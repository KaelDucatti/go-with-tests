package structs_methods_intefaces

import "errors"

const PI = 3.1415

func NewCircle(radius float32) (*Circle, error) {
	if radius <= 0 {
		return nil, errors.New("radius value cannot be equal to or less than 0")
	}

	return &Circle{Radius: radius}, nil
}

func (c *Circle) Perimeter() float32 {
	return PI * (2 * c.Radius)
}

func (c *Circle) Area() float32 {
	return PI * (c.Radius * c.Radius)
}