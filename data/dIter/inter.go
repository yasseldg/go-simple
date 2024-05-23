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
