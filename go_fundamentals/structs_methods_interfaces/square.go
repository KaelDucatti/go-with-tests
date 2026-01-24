package structs_methods_intefaces

import "errors"

func NewSquare(edge float32) (*Square, error) {
	if edge <= 0 {
		return nil, errors.New("value must be equal to or greater than 0")
	}

	return &Square{Edge: edge}, nil
}

func (s *Square) Perimeter() float32 {
	return s.Edge * 4
}

func (s *Square) Area() float32 {
	return s.Edge * s.Edge
}