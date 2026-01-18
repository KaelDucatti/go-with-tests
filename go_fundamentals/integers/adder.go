package integers

import "errors"

// Adder adds two integers and returns their sum.
// It returns an error if either number is zero.
func Adder(num1, num2 int) (int, error) {
	if num1 == 0 {
		return 0, errors.New("Num1 is required (cannot be zero).")
	} 
	if num2 == 0 {
		return 0, errors.New("Num2 is required (cannot be zero).")
	}

	return num1 + num2, nil
}