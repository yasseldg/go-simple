package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

type newCollFunc func(tSymbol.Inter, tInterval.Inter) (rMongo.InterRepo, error)

type repo struct {
	collFunc newCollFunc
}

func NewRepo(collFunc newCollFunc) InterRepo {
	return &repo{
		collFunc: collFunc,
	}
}

func (r *repo) GetIter(symbol tSymbol.Inter, interval tInterval.Inter) (InterIter, error) {
	return r.GetIterWithFilter(symbol, interval, nil)
}

func (r *repo) GetIterWithFilter(symbol tSymbol.Inter, interval tInterval.Inter, filter rFilter.Inter) (InterIter, error) {

	coll, err := r.collFunc(symbol, interval)
	if err != nil {
		return nil, fmt.Errorf("CandleCollFunc(): %s", err)
	}

	coll.Limit(2000)

	return NewIter(coll, filter), nil
}
