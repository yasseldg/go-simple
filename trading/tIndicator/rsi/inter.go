package rsi

import "github.com/yasseldg/go-simple/trading/tCandle"

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

// private

type inter interface {
	String() string
	Log()

	Periods() int

	Filled() bool
	Get() float64
}
