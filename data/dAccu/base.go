package dAccu

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Base struct {
	limit int

	count int

	empty bool
	err   error

	save func() error
}

func New(limit int, save func() error) *Base {
	return &Base{
		limit: limit,
		save:  save,
	}
}

func (a *Base) String(name string) string {
	return fmt.Sprintf("Accu ( %s ): count: %d  ..  limit: %d", name, a.count, a.limit)
}

func (a *Base) Log(name string) {
	sLog.Info(a.String(name))
}

func (a *Base) SetError(e error) {
	a.err = e
}

func (a Base) Error() error {
	return a.err
}

func (a *Base) Empty() bool {
	return a.empty
}

func (a *Base) Increase() {
	a.count++
}

func (a *Base) Limit() int {
	return a.limit
}

func (a *Base) Count() int {
	return a.count
}

func (a *Base) Save() {
	err := a.save()
	if err != nil {
		a.SetError(err)
		return
	}

	a.empty = true
}
