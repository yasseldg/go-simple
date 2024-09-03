package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

// Base
type Base struct {
	dIter.Inter

	coll   rMongo.InterColl
	filter rFilter.Inter
}

func New(filter rFilter.Inter, coll rMongo.InterColl) *Base {
	return &Base{
		Inter:  dIter.New(),
		coll:   coll,
		filter: filter,
	}
}

func (i *Base) String(name string) string {
	return fmt.Sprintf("%s coll: %s  ..  filter: %s", i.Inter.String(name), i.coll.String(), i.filter.Oper().String())
}

func (i *Base) Log(name string) {
	sLog.Info(i.String(name))
}

func (i *Base) Coll() rMongo.InterColl {
	return i.coll
}

func (i *Base) Filter() rFilter.Inter {
	return i.filter.Clone()
}