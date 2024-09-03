package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
)

type InterIter interface {
	rIter.Inter

	Item() Inter
	SetTsFrom(int64)
	SetTsTo(int64)
}

type Iter struct {
	rIter.Inter

	ts_from int64
	ts_to   int64

	item  Inter
	items Candles
}

func NewIter(filter rFilter.Inter, coll rMongo.InterColl) (*Iter, error) {

	return &Iter{Inter: rIter.New(filter, coll)}, nil
}

func (iter *Iter) String(name string) string {
	return fmt.Sprintf("%s ts_from: %s  ..  ts_to: %s", iter.Inter.String(name), sTime.ForLog(iter.ts_from, 0), sTime.ForLog(iter.ts_to, 0))
}

func (iter *Iter) Log(name string) {
	sLog.Info(iter.String(name))
}

func (iter *Iter) Item() Inter {
	return iter.item
}

func (iter *Iter) Next() bool {
	if !iter.Inter.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	filter := iter.Filter()
	filter.Ts(iter.ts_from, iter.ts_to)

	// sLog.Warn("next: filter: %v", filter)

	var items Candles
	err := iter.Coll().Filters(filter).Find(&items)
	if err != nil {
		iter.SetError(fmt.Errorf("next: coll.Find: %s", err))
		return false
	}

	if len(items) == 0 {
		iter.SetEmpty(true)
		return false
	}

	// sLog.Warn("next: items: %d", items[0].Ts())

	iter.items = items
	iter.ts_from = items[len(items)-1].Ts() + 1

	return iter.Next()
}

func (iter *Iter) SetTsFrom(ts_from int64) {
	iter.ts_from = ts_from
}

func (iter *Iter) SetTsTo(ts_to int64) {
	iter.ts_to = ts_to
}
