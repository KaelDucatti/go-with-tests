package array_and_slices

import "errors"

func Sum(slice []int) int {
	var total int
	for _, value := range slice {
		total += value
	}
	return total
}

func SumAll(slices ...[]int) ([]int, error) {
	if len(slices) == 0 {
		return nil, errors.New("SumAll must received at least 1 list")
	}
	
	for _, slice := range slices {
		if len(slice) == 0 {
			return nil, errors.New("all lists must have at least 1 value")
		}
	}

	var total []int
	for _, slice := range slices {
		total = append(total, Sum(slice))
	}

	return total, nil
}