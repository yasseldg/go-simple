package rAccu

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dAccu"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rMongo"

	"github.com/yasseldg/mgm/v4"
)

type Base struct {
	dAccu.Inter

	coll rMongo.Collection

	items []mgm.Model
}

func New(coll rMongo.Collection, limit int) *Base {

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
	sLog.Info(a.String(name))
}

func (a *Base) Clone() Inter {
	return New(a.coll, a.Limit())
}

func (a *Base) Add(model mgm.Model) {
	if model == nil {
		return
	}

	a.items = append(a.items, model)
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
		return fmt.Errorf("rMongo.CreateMany(): %s", err)
	}

	a.items = []mgm.Model{}

	return nil
}
