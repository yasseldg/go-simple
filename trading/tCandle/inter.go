package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rAccu"
	"github.com/yasseldg/go-simple/repos/rIter"
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
