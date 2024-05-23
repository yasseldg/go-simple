package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
)

// Base
type Base struct {
	dIter.Inter

	coll   rMongo.Collection
	filter rFilter.Filters
}

func New(filter rFilter.Filters, coll rMongo.Collection) *Base {
	return &Base{
		Inter:  dIter.New(),
		coll:   coll,
		filter: filter,
	}
}

func (i *Base) String(name string) string {
	return fmt.Sprintf("%s coll: %s  ..  filter: %s", i.Inter.String(name), i.coll.String(), i.filter.String())
}

func (i *Base) Log(name string) {
	sLog.Info(i.String(name))
}

func (i *Base) Coll() *rMongo.Collection {
	return &i.coll
}

func (i *Base) Filter() rFilter.Filters {
	return *i.filter.Clone()
}
