package tSymbol

import (
	"github.com/yasseldg/go-simple/data/dIter"
)

func NewIterLimited() InterIterLimited {
	return dIter.NewLimited[Inter](cloneFunc)
}

func cloneFunc(inter Inter) Inter {
	return inter.Clone()
}
