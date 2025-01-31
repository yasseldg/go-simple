package sFloats

import (
	"math"
	"slices"
)

// MinMax positive values of the float slice, including "0" or not
func MinMax(values []float64, zero bool) (float64, float64) {

	min := math.Inf(1)
	max := math.Inf(-1)

	for _, v := range values {
		if v > max {
			max = v
		}

		if v < min {
			if !zero && v == 0.0 {
				continue
			}
			min = v
		}
	}

	return min, max
}

// Sort the float slice in ascending or descending order
func Sort(values []float64, asc bool) {
	if asc {
		slices.Sort(values)
		return
	}

	slices.SortFunc(values, func(a, b float64) int {
		if a > b {
			return -1
		}
		if a < b {
			return 1
		}
		return 0
	})
}

func CompareSlices(s1, s2 []float64, prec int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if ComparePrec(s1[i], s2[i], prec) != 0 {
			return false
		}
	}

	return true
}

func SumSlice(s []float64) float64 {
	var sum float64
	for i := range s {
		sum += s[i]
	}
	return sum
}
