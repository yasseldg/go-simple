package priceStop

import "github.com/yasseldg/go-simple/trading/tCandle"

type InterPriceStop interface {
	String() string
	Log()

	SetStop(stop float64)
	Stop() float64
	Reset()
	Prev() tCandle.Inter

	Add(candle tCandle.Inter)
}
