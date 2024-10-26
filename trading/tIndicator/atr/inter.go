package atr

import "github.com/yasseldg/go-simple/trading/tCandle"

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
