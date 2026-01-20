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
	
	var totals []int
	for _, slice := range slices {
		if len(slice) == 0 {
			totals = append(totals, 0)
		} else {
			totals = append(totals, Sum(slice))
		}
	}

	return totals, nil
}

func SumAllTails(slices ...[]int) ([]int, error) {
	if len(slices) == 0 {
		return nil, errors.New("SumAllTails must received at least 1 list")
	}

	var totals []int
	for _, slice := range slices {
		if len(slice) == 0 {
			totals = append(totals, 0)
		} else {
			tail := slice[1:]
			totals = append(totals, Sum(tail))
		}
	}

	return totals, nil
}