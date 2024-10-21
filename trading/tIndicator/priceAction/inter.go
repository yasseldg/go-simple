package priceAction

import (
	"github.com/yasseldg/go-simple/trading/tIndicator/indicator"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sInts"
)

type Inter interface {
	indicator.Inter

	String() string
	Log()

	Add(candle tCandle.Inter)

	State() state
	IsTrigger() bool
	IsChange() bool

	Values() []ind

	InterPrices
	InterConfigs
}

type InterPrices interface {
	UpPrice() float64
	DownPrice() float64
	NeutralPrice() float64
}

type InterConfigs interface {
	SetSwing(bool) Inter
	SetHighLow(bool) Inter
	SetCloses(bool) Inter

	Swing() bool
	HighLow() bool
	Closes() bool

	ConfigNumber() int
}

type InterIterConfig interface {
	dIter.InterIterConfig

	Get() Inter

	SetSwing(sInts.InterIter) InterIterConfig
	SetHighLow(sInts.InterIter) InterIterConfig
	SetCloses(sInts.InterIter) InterIterConfig

	Swing() bool
	HighLow() bool
	Closes() bool

	Clone() InterIterConfig
}
