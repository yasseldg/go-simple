package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
	"github.com/yasseldg/go-simple/types/sTime"
)

type IterTs struct {
	Inter

	ts_from int64
	ts_to   int64
}

func NewTs(coll rMongo.InterColl, filter rFilter.Inter, sort rSort.Inter) *IterTs {

	if sort == nil {
		sort = rMongo.NewSort().TsAsc()
	}

	coll.Sorts(sort)
	coll.Limit(500)

	return &IterTs{
		Inter: New(coll, filter),
	}
}

func (iter *IterTs) String(name string) string {
	return fmt.Sprintf("%s ts_from: %s  ..  ts_to: %s",
		iter.Inter.String(name), sTime.ForLog(iter.ts_from, 0), sTime.ForLog(iter.ts_to, 0))
}

func (iter *IterTs) Log(name string) {
	sLog.Info(iter.String(name))
}

func (iter *IterTs) TsFrom() int64 {
	return iter.ts_from
}

func (iter *IterTs) TsTo() int64 {
	return iter.ts_to
}

func (iter *IterTs) SetTsFrom(ts_from int64) {
	iter.ts_from = ts_from
}

func (iter *IterTs) SetTsTo(ts_to int64) {
	iter.ts_to = ts_to
}
