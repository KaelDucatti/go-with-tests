package structs_methods_intefaces

type Form interface {
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

func Area(f Form) float32 {
	return f.Area()
}

func Perimeter(f Form) float32 {
	return f.Perimeter()
}