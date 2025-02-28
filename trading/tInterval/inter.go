package tInterval

import "github.com/yasseldg/go-simple/data/dIter"

type Inter interface {
	String() string

	Get() Interval
	AsaName(int64) string

	ForLog() string
	IsDefault() bool
	IsClosing(int64, Inter) bool

	Minutes() int64
	Seconds() int64
	MilliSeconds() int64

	Prev(int64) int64
	Next(int64) int64
}

type Inters []Inter

// InterIterLimited
type InterIterLimited interface {
	dIter.InterLimited[Inter]
}
