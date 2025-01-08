package dIter

import "github.com/yasseldg/go-simple/logs/sLog"

type Inter interface {
	sLog.InterStringLogName

	SetError(error)
	Error() error

	SetEmpty(bool)
	Empty() bool

	Next() bool
}

// Define a generic interface
type InterLimited[T any] interface {
	Inter

	Add(...T)
	Reset()
	Item() T
	Count() int

	Clone() InterLimited[T]
}

// Config interfaces

type InterConfig interface {
	Inter

	Count() int
	Reset()
}

type InterNameConfig interface {
	InterConfig
	Name() string
}

type InterIterConfig interface {
	InterNameConfig

	Add(...InterNameConfig)
}
