package sStrings

import "strings"

func ToUpper(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}

	uppers := make([]string, len(s))
	for i, v := range s {
		if len(v) == 0 {
			continue
		}
		uppers[i] = strings.ToUpper(v)
	}

	return uppers
}
