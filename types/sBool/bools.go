package sBool

import "strconv"

func Get(str string) bool {
	v, _ := strconv.ParseBool(str)

	return v
}

func ToString(v bool) string {
	return strconv.FormatBool(v)
}
