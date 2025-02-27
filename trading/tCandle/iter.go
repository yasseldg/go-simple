package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Iter struct {
	*rIter.ModelTs[*model]
}

func NewIter(coll rMongo.InterRepo, filter rFilter.Inter) *Iter {
	return &Iter{
		ModelTs: rIter.NewModelTs[*model](coll, filter, nil),
	}
}

func (iter *Iter) Item() Inter {
	return iter.ModelTs.Item
}

func (iter *Iter) Clone() InterIter {
	return NewIter(iter.Coll().Clone(), iter.Filter().Clone())
}
