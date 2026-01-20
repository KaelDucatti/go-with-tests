package array_and_slices

import "errors"

func Sum(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("the list must have at least 1 value")
	}

	var total int
	for _, value := range slice {
		total += value
	}
	return total, nil
}

func SumAll(matrix [][]int) ([]int, error) {
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) <= 0 {
			return nil, errors.New("all lists must have at least 1 value")
		}
	}
	
	total := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			total[i] += matrix[i][j]
		} 
	}

	return total, nil
}