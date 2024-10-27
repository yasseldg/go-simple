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
func NewADX(period int) adx.Inter {
	return adx.New(period)
}

// ATR
func NewAvgATR(period int) atr.Inter {
	return atr.NewAvg(period)
}

func NewSmoothedATR(period int) atr.Inter {
	return atr.NewSmoothed(period)
}

// RSI
func NewRSI(period int) rsi.Inter {
	return rsi.New(period)
}

func NewRSIcandle(period int) rsi.InterCandle {
	return rsi.NewCandle(period)
}

// BollingerBands
func NewBB(periods int, deviations float64) bollingerBands.Inter {
	return bollingerBands.New(periods, deviations)
}

func NewBBcandle(periods int, deviations float64) bollingerBands.InterCandle {
	return bollingerBands.NewCandle(periods, deviations)
}

// Price Action
func NewPriceAction() priceAction.Inter {
	return priceAction.New()
}

// Super Trend
func NewSuperTrend(periods int, multiplier float64, smoothed bool) superTrend.Inter {
	return superTrend.New(periods, multiplier, smoothed)
}

//  Iter Configs

func NewATRIterConfig(name string) atr.InterIterConfig {
	return atr.NewIterConfig(name)
}

func NewBBIterConfig(name string) bollingerBands.InterIterConfig {
	return bollingerBands.NewIterConfig(name)
}

func NewPriceActionIterConfig(name string) priceAction.InterIterConfig {
	return priceAction.NewIterConfig(name)
}

func NewSuperTrendIterConfig(name string) superTrend.InterIterConfig {
	return superTrend.NewIterConfig(name)
}
