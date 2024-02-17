package sFloats

import (
	"math"
	"strconv"
)

func Get64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)

	return GetValid(v)
}

func ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func GetValid(f float64) float64 {
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
