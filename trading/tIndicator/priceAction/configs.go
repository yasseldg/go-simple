package priceAction

import (
	"github.com/yasseldg/go-simple/trading/tCandle"
)

func (pa *base) SetSwing(v bool) Inter {
	pa.swing = v
	return pa
}

func (pa *base) SetHighLow(v bool) Inter {
	if v && pa.closes {
		return pa
	}

	pa.high_low = v
	return pa
}

func (pa *base) SetCloses(v bool) Inter {
	pa.closes = v

	if v {
		pa.funcHigh = func(candle tCandle.Inter) float64 { return candle.Close() }
		pa.funcLow = func(candle tCandle.Inter) float64 { return candle.Close() }
		pa.SetHighLow(false)
	} else {
		pa.funcHigh = func(candle tCandle.Inter) float64 { return candle.High() }
		pa.funcLow = func(candle tCandle.Inter) float64 { return candle.Low() }
	}

	return pa
}

func (pa *base) Swing() bool {
	return pa.swing
}

func (pa *base) HighLow() bool {
	return pa.high_low
}

func (pa *base) Closes() bool {
	return pa.closes
}

func (pa *base) ConfigNumber() int {
	switch {
	case !pa.HighLow() && !pa.Swing() && !pa.Closes():
		return 2

	case pa.HighLow() && !pa.Swing() && !pa.Closes():
		return 3

	case !pa.HighLow() && pa.Swing() && !pa.Closes():
		return 4

	case pa.HighLow() && pa.Swing() && !pa.Closes():
		return 5

	case !pa.HighLow() && !pa.Swing() && pa.Closes():
		return 6

	default:
		return 0
	}
}
