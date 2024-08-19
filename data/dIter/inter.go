package dIter

type Inter interface {
	String(name string) string
	Log(name string)

	SetError(e error)
	Error() error

	SetEmpty(e bool)
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

	Add(config ...InterNameConfig)
}
