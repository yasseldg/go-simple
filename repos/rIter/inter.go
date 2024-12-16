package rIter

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Inter interface {
	dIter.Inter

	Coll() rMongo.InterRepo
	Filter() rFilter.Inter
}

type InterTs interface {
	Inter

	TsFrom() int64
	TsTo() int64

	SetTsFrom(int64)
	SetTsTo(int64)
}
