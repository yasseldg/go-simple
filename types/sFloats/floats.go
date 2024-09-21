package sFloats

import (
	"math"
	"strconv"
)

func Get64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)

	return GetValid(v)
}

func Sum(values []float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func ToString(v float64, prec int) string {
	return strconv.FormatFloat(v, 'f', prec, 64)
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

func GetPrec(value float64, prec int) float64 {
	pow := math.Pow10(prec)
	return math.Round(value*pow) / pow
}

func GetTrunc(value float64, prec int) float64 {
	pow := math.Pow10(prec)
	return float64(int64(value*pow)) / pow
}

// trunc a float value an add a plus value at prec decimal
func GetTruncPlus(value float64, prec, plus int) float64 {
	pow := math.Pow10(prec)
	truncated := float64(int64(value*pow)) / pow
	return truncated + float64(plus)/pow
}

// ComparePrec, "v1>v2: >0",  "v1<v2: <0",  "v1=v2: 0"
func ComparePrec(v1, v2 float64, prec int) float64 {
	return GetPrec(v1, prec) - GetPrec(v2, prec)
}

// CompareTrunc, "v1>v2: >0",  "v1<v2: <0",  "v1=v2: 0"
func CompareTrunc(v1, v2 float64, prec int) float64 {
	return GetTrunc(v1, prec) - GetTrunc(v2, prec)
}

func Compare(v1, v2 float64, prec int) int {
	epsilon := 1.0 / math.Pow(10.0, float64(prec))

	diff := v1 - v2
	if math.Abs(diff) < epsilon {
		return 0 // v1 == v2 (within precision)
	}
	if diff > 0 {
		return 1 // v1 > v2
	}
	return -1 // v1 < v2
}

func GetDiffPercent(fromValue, toValue float64) float64 {

	return ((toValue / fromValue) - 1) * 100
}

func GetDiffPercentAbs(fromValue, toValue float64) float64 {

	return math.Abs(GetDiffPercent(fromValue, toValue))
}

func GetDiffAbs(fromValue, toValue float64) float64 {

	return math.Abs(fromValue - toValue)
}
