package sStrings

import "github.com/yasseldg/go-simple/data/dIter"

func NewIterLimited(values ...string) InterIterLimited {

	iter := dIter.NewLimited[string](nil)
	iter.Add(values...)

	return iter
}
