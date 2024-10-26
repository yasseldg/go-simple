package bollingerBands

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

//  BBcandle is a BBands indicator for candles

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

	SetPeriods(sInts.InterIter)
	SetDeviations(sFloats.InterIter)

	Get() Inter
	GetCandle() InterCandle
	Clone() InterIterConfig
}

// private

type inter interface {
	String() string
	Log()

	Periods() int
	Deviations() float64

	Filled() bool
	Get() (mean, upper, lower float64)
	Calc(deviations float64) (mean, upper, lower float64)
}
