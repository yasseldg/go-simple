package atr

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sInts"
)

// ATR (Average True Range)

type Inter interface {
	String() string
	Log()

	Count() int
	Periods() int
	Filled() bool
	Get() float64
	Prev() tCandle.Inter

	Add(candle tCandle.Inter)
}

type InterIterConfig interface {
	dIter.InterIterConfig

	SetPeriods(sInts.InterIter)
	SetSmoothed(sInts.InterIter)

	Smoothed() bool

	Get() Inter

	Clone() InterIterConfig
}
