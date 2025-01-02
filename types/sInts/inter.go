package sInts

import "github.com/yasseldg/go-simple/data/dIter"

type InterIter interface {
	dIter.Inter

	Value() int64
	Count() int
	Reset()

	Clone() InterIter
}
