package rIter

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
	"github.com/yasseldg/go-simple/trading/tInterval"
)

type IntervalTs[T any] struct {
	InterTs

	mu sync.Mutex

	interval tInterval.Inter
	limit    int64

	item  *T
	items []*T

	getItems FuncGetItems[T]
}

type FuncGetItems[T any] func(rMongo.D) ([]*T, int64)

func NewIntervalTs[T any](coll rMongo.InterRepo, filter rFilter.Inter, sort rSort.Inter,
	interval tInterval.Inter, limit int, getItems FuncGetItems[T]) *IntervalTs[T] {
	it := &IntervalTs[T]{
		InterTs:  NewTs(coll, filter, sort),
		interval: interval,
		limit:    int64(limit),
		getItems: getItems,
	}
	it.SetTsFrom(0)

	return it
}

func (it *IntervalTs[T]) Limit() int64 {
	return it.limit
}

func (it *IntervalTs[T]) Interval() tInterval.Inter {
	return it.interval
}

func (it *IntervalTs[T]) String(name string) string {
	return fmt.Sprintf("%s  ..  interval: %s  ..  limit: %d",
		it.InterTs.String(name), it.interval.String(), it.limit)
}

func (it *IntervalTs[T]) SetTsFrom(ts int64) {
	it.mu.Lock()
	defer it.mu.Unlock()

	it.setTsFrom(ts)
}

func (it *IntervalTs[T]) setTsFrom(ts int64) {
	if ts <= 0 {
		ts = it.Coll().FirstTs(0, 0)
	}

	it.InterTs.SetTsFrom(ts)
}

func (it *IntervalTs[T]) Item() *T {
	it.mu.Lock()
	defer it.mu.Unlock()

	return it.item
}

func (it *IntervalTs[T]) Next() bool {
	it.mu.Lock()
	defer it.mu.Unlock()

	return it.next()
}

func (it *IntervalTs[T]) next() bool {

	if !it.InterTs.Next() {
		return false
	}

	if len(it.items) > 0 {
		it.item = it.items[0]
		it.items = it.items[1:]
		return true
	}

	if it.TsFrom() <= 0 {
		it.SetEmpty(true)
		return false
	}

	ts_to := it.TsFrom() + (it.limit * it.interval.Seconds())
	if it.TsTo() > 0 && ts_to > it.TsTo() {
		ts_to = it.TsTo()
	}

	filter := it.Filter().Ts(it.TsFrom(), ts_to)

	ts_filter, err := rMongo.FilterFields(filter)
	if err != nil {
		it.SetError(fmt.Errorf("iter: rMongo.FilterFields(): %s", err))
		return false
	}

	items, ts_from := it.getItems(ts_filter)

	if it.Empty() {
		return false
	}

	it.items = items
	it.setTsFrom(ts_from)

	return it.next()
}
