package sInts

import "strconv"

func Get(str string) int {

	v, _ := strconv.Atoi(str)

	return v
}

func Get64(str string) int64 {

	v, _ := strconv.ParseInt(str, 10, 64)

	return v
}
