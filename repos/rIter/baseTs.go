package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
	"github.com/yasseldg/go-simple/types/sTime"
)

type iterTs struct {
	Inter

	ts_from int64
	ts_to   int64
}

// by default coll.limit = 500
func NewTs(coll rMongo.InterRepo, filter rFilter.Inter, sort rSort.Inter) *iterTs {

	if sort == nil {
		sort = rMongo.NewSort().TsAsc()
	}

	coll.Limit(500)

	return &iterTs{
		Inter: New(coll, filter, sort),
	}
}

func (it *iterTs) String(name string) string {
	return fmt.Sprintf("%s ts_from: %s  ..  ts_to: %s", it.Inter.String(name),
		sTime.ForLog(it.ts_from, 0), sTime.ForLog(it.ts_to, 0))
}

func (it *iterTs) Log(name string) {
	sLog.Info(it.String(name))
}

func (it *iterTs) TsFrom() int64 {
	return it.ts_from
}

func (it *iterTs) TsTo() int64 {
	return it.ts_to
}

func (it *iterTs) SetTsFrom(ts_from int64) {
	it.ts_from = ts_from
}

func (it *iterTs) SetTsTo(ts_to int64) {
	it.ts_to = ts_to
}
