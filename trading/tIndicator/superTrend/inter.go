package superTrend

import (
	"github.com/yasseldg/go-simple/trading/tIndicator/indicator"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
)

type Inter interface {
	indicator.Inter

	String() string
	Log()

	Add(candle tCandle.Inter)

	IsUptrend() bool
	IsDowntrend() bool
}

type InterIterConfig interface {
	dIter.InterIterConfig

	Get() Inter

	Clone() InterIterConfig
}
