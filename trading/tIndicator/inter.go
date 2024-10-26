package tIndicator

import (
	"github.com/yasseldg/go-simple/trading/tIndicator/adx"
	"github.com/yasseldg/go-simple/trading/tIndicator/atr"
	"github.com/yasseldg/go-simple/trading/tIndicator/bollingerBands"
	"github.com/yasseldg/go-simple/trading/tIndicator/priceAction"
	"github.com/yasseldg/go-simple/trading/tIndicator/rsi"
	"github.com/yasseldg/go-simple/trading/tIndicator/superTrend"
)

// ADX
type InterADX interface {
	adx.Inter
}

// ATR
type InterATR interface {
	atr.Inter
}

// RSI
type InterRSI interface {
	rsi.Inter
}

type InterRSIcandle interface {
	rsi.InterCandle
}

// BollingerBands
type InterBBands interface {
	bollingerBands.Inter
}

type InterBBandsCandle interface {
	bollingerBands.InterCandle
}

// Super Trend
type InterSuperTrend interface {
	superTrend.Inter
}

// Price Action
type InterPriceAction interface {
	priceAction.Inter
}

//  Iter Configs

type InterADXIterConfig interface {
}

type InterATRIterConfig interface {
}

type InterRSIIterConfig interface {
}

type InterBBandsIterConfig interface {
	bollingerBands.InterIterConfig
}

type InterPriceActionIterConfig interface {
	priceAction.InterIterConfig
}

type InterSuperTrendIterConfig interface {
	superTrend.InterIterConfig
}
