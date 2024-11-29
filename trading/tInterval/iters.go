package tInterval

import (
	"github.com/yasseldg/go-simple/data/dIter"
)

func NewIterLimited(intervals ...Inter) InterIterLimited {

	iter := dIter.NewLimited[Inter](nil)
	iter.Add(intervals...)

	return iter
}
