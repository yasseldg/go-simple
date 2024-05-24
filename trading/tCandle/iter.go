package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rIter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
)

type InterIter interface {
	rIter.Inter

	Item() Inter
	SetNextTs(int64)
}

type Iter struct {
	rIter.Inter

	next_ts int64

	item  Inter
	items Candles
}

func NewIter(filter rFilter.Filters, coll rMongo.Collection) (*Iter, error) {

	return &Iter{Inter: rIter.New(filter, coll)}, nil
}

func (iter *Iter) String(name string) string {
	return fmt.Sprintf("%s next_ts: %d", iter.Inter.String(name), iter.next_ts)
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
	filter.Ts(iter.next_ts, 0)

	// sLog.Debug("next: filter: %v", filter)

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

	// sLog.Warn("next: items: %d", items[0].UnixTs)

	iter.items = items
	iter.next_ts = items[len(items)-1].Ts() + 1

	return iter.Next()
}

func (iter *Iter) SetNextTs(next int64) {
	iter.next_ts = next
}
