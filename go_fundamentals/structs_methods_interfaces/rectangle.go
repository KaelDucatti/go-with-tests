package structs_methods_intefaces

const (
	ErrValueIsZero             = RectangleErr("You cannot input the number 0")
	ErrValueIsNegative         = RectangleErr("You cannot input a negative value")
	ErrHeightAndWidthAreEquals = RectangleErr("In a Rectangle Height and Width cannot be equals.")
)

type RectangleErr string

func (re RectangleErr) Error() string {
	return string(re)
}

func NewRectangle(height, width float32) (*Rectangle, error) {
	if err := validateIfValueIsNegative(height, width); err != nil {
		return nil, err
	}
	if err := validateIfValueIsZero(height, width); err != nil {
		return nil, err
	}
	if err := validateIfValuesAreEquals(height, width); err != nil {
		return nil, err
	}	
	return &Rectangle{
		height: height,
		width:  width,
	}, nil
}

func (r *Rectangle) Area() float32 {
	return r.height * r.width
}

func (r *Rectangle) Perimeter() float32 {
	return 2 * (r.height + r.width)
}

func (r *Rectangle) Height() float32 {
	return r.height
}

func (r *Rectangle) Width() float32 {
	return r.width
}

func (r *Rectangle) SetHeight(value float32) error {
	if err := validateIfValueIsNegative(value); err != nil {
		return err
	}
	if err := validateIfValueIsZero(value); err != nil {
		return err
	}
	if err := validateIfValuesAreEquals(value, r.Width()); err != nil {
		return err
	}
	r.height = value
	return nil
}

func (r *Rectangle) SetWidth(value float32) error {
	if err := validateIfValueIsNegative(value); err != nil {
		return err
	}
	if err := validateIfValueIsZero(value); err != nil {
		return err
	}
	if err := validateIfValuesAreEquals(value, r.Height()); err != nil {
		return err
	}
	r.width = value
	return nil
}

func validateIfValueIsZero(values ...float32) error {
	for _, value := range values {
		if value == 0 {
			return ErrValueIsZero
		}
	}
	return nil
}

func validateIfValueIsNegative(values ...float32) error {
	for _, value := range values {
		if value < 0 {
			return ErrValueIsNegative
		}
	}
	return nil
}

func validateIfValuesAreEquals(height, width float32) error {
	if height == width {
		return ErrHeightAndWidthAreEquals
	}
	return nil
}
