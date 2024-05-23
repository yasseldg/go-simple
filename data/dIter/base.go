package dIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Base struct {
	empty bool
	err   error
}

func New() *Base {
	return &Base{}
}

func (a *Base) String(name string) string {
	return fmt.Sprintf("Iter ( %s ): ", name)
}

func (a *Base) Log(name string) {
	sLog.Info(a.String(name))
}

func (b *Base) SetError(e error) {
	b.err = e
}

func (b *Base) Error() error {
	return b.err
}

func (b *Base) SetEmpty(e bool) {
	b.empty = e
}

func (b *Base) Empty() bool {
	return b.empty
}

func (b *Base) Next() bool {
	if b.empty {
		return false
	}

	if b.err != nil {
		return false
	}

	return true
}
