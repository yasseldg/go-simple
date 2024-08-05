package tInterval

import (
	"github.com/yasseldg/go-simple/data/dIter"
)

func NewIterLimited() InterIterLimited {
	return dIter.NewLimited[Inter](nil)
}
