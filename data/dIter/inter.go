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

type InterConfig interface {
	Inter

	Count() int
	Reset()
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
