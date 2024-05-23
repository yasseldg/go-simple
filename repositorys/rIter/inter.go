package rIter

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/repositorys/rFilter"
	"github.com/yasseldg/go-simple/repositorys/rMongo"
)

type Inter interface {
	dIter.Inter

	Coll() *rMongo.Collection
	Filter() rFilter.Filters
}
