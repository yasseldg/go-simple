package dAccu

import "github.com/yasseldg/go-simple/data/dTs"

type Inter interface {
	String(string) string
	Log(string)

	SetError(error)
	Error() error

	Increase()
	Limit() int
	Count() int
	Empty() bool

	Save()
}

type InterTs interface {
	Inter

	Add(dTs.Inter)
}
