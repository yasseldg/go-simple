package rAccu

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/data/dAccu"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Base struct {
	dAccu.Inter

	coll rMongo.InterColl

	mu sync.Mutex

	items []rMongo.InterModel

	errs int
}

func New(coll rMongo.InterColl, limit int) *Base {

	b := &Base{
		coll: coll,
	}

	b.Inter = dAccu.New(limit, b.save)

	return b
}

func (a *Base) String(name string) string {
	return fmt.Sprintf("%s  ..  items: %d  ..  %s", a.Inter.String(name), len(a.items), a.coll.String())
}

func (a *Base) Log(name string) {
	if a.Error() != nil {
		sLog.Error("%s  ..  err: %s", a.String(name), a.Error())
		return
	}
	sLog.Info(a.String(name))
}

func (a *Base) Clone() Inter {
	return New(a.coll.Clone(), a.Limit())
}

func (a *Base) Add(inter rMongo.InterModel) {
	if inter == nil {
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	a.items = append(a.items, inter)
	a.Increase()

	if len(a.items) >= a.Limit() {
		a.Save()
	}
}

//  private methods

func (a *Base) save() error {
	if len(a.items) == 0 {
		return nil
	}

	err := rMongo.CreateMany(a.items, a.coll)
	if err != nil {
		return a.saveErr()
	}

	a.items = []rMongo.InterModel{}

	return nil
}

func (a *Base) saveErr() error {
	c := 0
	for _, item := range a.items {
		err := a.coll.Create(item)
		if err != nil {
			continue
		}
		c++
	}
	a.errs += len(a.items) - c

	a.items = []rMongo.InterModel{}

	return fmt.Errorf("rMongo.CreateMany(): %d items", a.errs)
}
