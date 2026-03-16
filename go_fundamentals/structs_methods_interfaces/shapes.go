package structs_methods_intefaces

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Rectangle struct {
	height float32
	width  float32
}

type Square struct {
	Edge float32
}

type Circle struct {
	Radius float32
}
