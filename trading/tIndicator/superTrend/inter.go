package superTrend

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tIndicator/atr"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

// Super Trend like TradingView

type Inter interface {
	atr.Inter

	Add(candle tCandle.Inter)

	Multiplier() float64
	Config() string

	IsUptrend() bool
	IsDowntrend() bool
}

type InterIterConfig interface {
	dIter.InterIterConfig

	SetPeriods(sInts.InterIter)
	SetMultiplier(sFloats.InterIter)
	SetSmoothed(sInts.InterIter)

	Get() (Inter, error)
	Clone() InterIterConfig
}
