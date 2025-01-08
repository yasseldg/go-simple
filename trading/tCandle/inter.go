package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rAccu"
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIter"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

type Inter interface {
	String(prec int) string
	Log(prec int)

	Ts() int64
	Open() float64
	High() float64
	Low() float64
	Close() float64
	Volume() float64

	InterCalc
}

type InterCalc interface {
	LogReturn() float64
	BodyPerc() float64
}

type InterModel interface {
	Inter
	GetModel() *Candle
}

type InterIter interface {
	rIter.InterTs

	Item() Inter

	Clone() InterIter
}

type InterAccu interface {
	rAccu.Inter

	AddCandle(*Candle)
}

type InterRepo interface {
	GetIter(tSymbol.Inter, tInterval.Inter) (InterIter, error)
	GetIterWithFilter(tSymbol.Inter, tInterval.Inter, rFilter.Inter) (InterIter, error)
}
