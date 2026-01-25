package structs_methods_intefaces

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