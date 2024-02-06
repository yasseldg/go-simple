package sFloats

import "math"

func Get64(f float64) float64 {
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return 0
	}
	return f
}

// Adding or Subtracting percent to value
func GetWithPercent(value, percent float64, adding bool) float64 {

	if adding {
		return ((percent / 100) + 1) * value
	}
	return -((percent / 100) - 1) * value
}
