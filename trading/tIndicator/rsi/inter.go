package rsi

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sInts"
)

// RSI (Relative Strength Index)

type Inter interface {
	inter

	Add(close float64)
}

type InterCandle interface {
	inter

	Add(candle tCandle.Inter)
	Candle() tCandle.Inter
}

type InterIterConfig interface {
	dIter.InterIterConfig

	Get() (Inter, error)

	SetPeriods(sInts.InterIter) InterIterConfig
	Periods() int64

	Clone() InterIterConfig
}

// private

type inter interface {
	String() string
	Log()

	Periods() int

	Filled() bool
	Get() float64
}
