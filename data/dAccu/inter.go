package dAccu

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
