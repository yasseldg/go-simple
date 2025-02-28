package rIter

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/repos/rSort"
)

type ModelTs[T rMongo.InterModelTs] struct {
	InterTs

	Item  T
	items []T
}

func NewModelTs[T rMongo.InterModelTs](coll rMongo.InterRepo, filter rFilter.Inter, sort rSort.Inter) *ModelTs[T] {
	return &ModelTs[T]{
		InterTs: NewTs(coll, filter, sort),
	}
}

func (it *ModelTs[T]) Next() bool {
	if !it.InterTs.Next() {
		return false
	}

	if len(it.items) > 0 {
		it.Item = it.items[0]
		it.items = it.items[1:]
		return true
	}

	filter := it.Filter().Ts(it.TsFrom(), it.TsTo())

	var items []T
	err := it.Coll().Filters(filter).Sorts(it.Sort()).Find(&items)
	if err != nil {
		it.SetError(fmt.Errorf("next: coll.Find: %s", err))
		return false
	}

	if len(items) == 0 {
		it.SetEmpty(true)
		return false
	}

	it.items = items
	it.SetTsFrom(items[len(items)-1].Ts() + 1)

	return it.Next()
}
