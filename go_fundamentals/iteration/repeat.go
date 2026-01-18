package iteration

import (
	"errors"
	"strings"
)

// Repeat repeats five times the entered value
// It returns error if value is empty
func Repeat(value string, count int) (string, error) {
	if value == "" {
		return "", errors.New("value cannot be empty")
	}
	if count <= 0 {
		return "", errors.New("count must be greater than 0")
	}

	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(value)
	}

	return repeated.String(), nil
}