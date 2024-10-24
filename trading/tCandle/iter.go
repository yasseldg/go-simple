package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Iter struct {
	rIter.InterTs

	item  Inter
	items Candles
}

func NewIter(coll rMongo.InterColl, filter rFilter.Inter) *Iter {
	return &Iter{
		InterTs: rIter.NewTs(coll, filter, nil),
	}
}

func (iter *Iter) Item() Inter {
	return iter.item
}

func (iter *Iter) Next() bool {
	if !iter.InterTs.Next() {
		return false
	}

	if len(iter.items) > 0 {
		iter.item = iter.items[0]
		iter.items = iter.items[1:]
		return true
	}

	filter := iter.Filter().Ts(iter.TsFrom(), iter.TsTo())

	var items Candles
	err := iter.Coll().Filters(filter).Find(&items)
	if err != nil {
		iter.SetError(fmt.Errorf("next: coll.Find: %s", err))
		sLog.Error(iter.Error().Error())
		return false
	}

	if len(items) == 0 {
		iter.SetEmpty(true)
		return false
	}

	iter.items = items
	iter.SetTsFrom(items[len(items)-1].Ts() + 1)

	return iter.Next()
}

func (iter *Iter) Clone() InterIter {
	return NewIter(iter.Coll().Clone(), iter.Filter().Clone())
}
