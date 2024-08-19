package sInts

import (
	"math"
	"strconv"
)

func Get(str string) int {

	v, _ := strconv.Atoi(str)

	return v
}

func Get64(str string) int64 {

	v, _ := strconv.ParseInt(str, 10, 64)

	return v
}

func ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func InflateFloat64(v float64, prec int) int64 {
	return int64(v * math.Pow10(prec))
}

func DeflateFloat64(v int64, prec int) float64 {
	return float64(v) / math.Pow10(prec)
}
