package rIter

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	dIter.Inter

	Coll() rMongo.InterColl
	Filter() rFilter.Inter
}
