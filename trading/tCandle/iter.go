package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rIter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
)

type Iter struct {
	rIter.Iter

	next_ts int64

	Item  Candle
	items Candles
}

func NewIter(filter rFilter.Filters, coll rMongo.Collection) (Iter, error) {

	return Iter{Iter: rIter.NewIter(filter, coll)}, nil
}

func (iter *Iter) Next() bool {
	if !iter.Iter.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.Item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	filter := iter.Filter()
	filter.Ts(iter.next_ts, 0)

	// sLog.Debug("next: filter: %v", filter)

	var items Candles
	err := iter.Coll.Filters(filter).Find(&items)
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
	iter.next_ts = items[len(items)-1].Ts + 1

	return iter.Next()
}

func (iter *Iter) SetNextTs(next int64) {
	iter.next_ts = next
}
