package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
)

// Base
type Base struct {
	dIter.Inter

	coll   rMongo.InterRepo
	filter rFilter.Inter
	sort   rSort.Inter
}

func New(coll rMongo.InterRepo, filter rFilter.Inter, sort rSort.Inter) *Base {

	if coll == nil {
		return nil
	}

	if filter == nil {
		filter = rMongo.NewFilter()
	}

	if sort == nil {
		sort = rMongo.NewSort()
	}

	return &Base{
		Inter:  dIter.New(),
		coll:   coll,
		filter: filter,
		sort:   sort,
	}
}

func (i *Base) String(name string) string {
	return fmt.Sprintf("%s %s  ..  filter: %s", i.Inter.String(name), i.coll.String(), i.filter.Oper().String())
}

func (i *Base) Log(name string) {
	sLog.Info(i.String(name))
}

func (i *Base) Coll() rMongo.InterRepo {
	return i.coll.Clone()
}

func (i *Base) Filter() rFilter.Inter {
	return i.filter.Clone()
}

func (i *Base) Sort() rSort.Inter {
	return i.sort.Clone()
}
