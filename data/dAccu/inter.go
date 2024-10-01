package dAccu

import "github.com/yasseldg/go-simple/data/dTs"

type Inter interface {
	String(name string) string
	Log(name string)

	SetError(e error)
	Error() error

	Increase()
	Limit() int
	Count() int
	Empty() bool

	Save()
}

type InterTs interface {
	Inter

	Add(inter dTs.Inter)
}
