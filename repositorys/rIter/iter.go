package rIter

import (
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
)

//  Iter

type Iter struct {
	Coll   rMongo.Collection
	filter rFilter.Filters

	empty bool
	err   error
}

func NewIter(filter rFilter.Filters, coll rMongo.Collection) Iter {
	return Iter{
		Coll:   coll,
		filter: filter,
	}
}

func (i Iter) Next() bool {
	if i.empty {
		return false
	}

	if i.err != nil {
		return false
	}

	return true
}

func (i *Iter) SetError(e error) {
	i.err = e
}

func (i Iter) Error() error {
	return i.err
}

func (i *Iter) SetEmpty(e bool) {
	i.empty = e
}

func (i Iter) Empty() bool {
	return i.empty
}

func (i Iter) Filter() rFilter.Filters {
	return *i.filter.Clone()
}
