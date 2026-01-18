package array_and_slices

import "errors"

func Sum(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("slice must have at least 1 value")
	}

	var total int
	for _, value := range slice {
		total += value
	}
	return total, nil
}