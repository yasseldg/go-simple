package sFloats

import "github.com/yasseldg/go-simple/data/dIter"

type InterIter interface {
	dIter.Inter

	Value() float64
	Count() int
	Reset()

	Clone() InterIter
}
