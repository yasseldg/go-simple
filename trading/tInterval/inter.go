package tInterval

import "github.com/yasseldg/go-simple/data/dIter"

type Inter interface {
	String() string

	ForLog() string
	IsDefault() bool
	IsClosing(int64) bool

	Minutes() int64
	Seconds() int64
	MilliSeconds() int64

	Prev(ts int64) int64
	Next(ts int64) int64
}

type Inters []Inter

// InterIterLimited
type InterIterLimited interface {
	dIter.InterLimited[Inter]
}
