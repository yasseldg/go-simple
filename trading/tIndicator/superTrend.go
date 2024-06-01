package tIndicator

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sStrings"
	"github.com/yasseldg/go-simple/types/sTime"
)

// SuperTrend like TradingView

type InterSuperTrend interface {
	InterATR

	String() string

	Config() string
	Multiplier() float64
	IsUptrend() bool
	Add(candle tCandle.Inter)
	Get() float64

	Periods() int
}

type SuperTrend struct {
	InterATR

	mu sync.Mutex

	multiplier float64

	prev_close float64
	prev_upper float64
	prev_value float64

	upper float64
	lower float64

	value float64
}

func NewSuperTrend(periods int, multiplier float64, smoothed bool) *SuperTrend {
	supertrend := &SuperTrend{
		multiplier: multiplier,
	}

	if smoothed {
		supertrend.InterATR = NewSmATR(periods)
	} else {
		supertrend.InterATR = NewAvgATR(periods)
	}

	return supertrend
}

func (st *SuperTrend) Config() string {
	return fmt.Sprintf("period: %d  ..  multiplier: %.2f", st.Periods(), st.multiplier)
}

func (st *SuperTrend) String() string {
	v := ""
	if st.IsUptrend() {
		v = sStrings.Colored(sStrings.Green, fmt.Sprintf("%f", st.value))
	} else {
		v = sStrings.Colored(sStrings.Red, fmt.Sprintf("%f", st.value))
	}

	return fmt.Sprintf("SuperTrend %d: period: %d  ..  multiplier: %.2f  ..  %s  ..  %s",
		st.InterATR.Count(), st.Periods(), st.multiplier, sTime.ForLog(st.InterATR.Prev().Ts(), 0), v)
}

func (st *SuperTrend) Log() {
	st.mu.Lock()
	defer st.mu.Unlock()

	sLog.Info(st.String())
}

func (st *SuperTrend) Multiplier() float64 {
	return st.multiplier
}

func (st *SuperTrend) IsUptrend() bool {
	return st.value == st.lower
}

func (st *SuperTrend) Add(candle tCandle.Inter) {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.InterATR.Add(candle)

	if st.InterATR.Get() == 0 {
		return
	}

	hl2 := (candle.High() + candle.Low()) / 2

	multAtr := st.multiplier * st.InterATR.Get()

	basicUpper := hl2 + multAtr
	basicLower := hl2 - multAtr

	// sLog.Debug("SuperTrend %d: basicUpper: %f  ..  uper: %f  ..  basicLower: %f  ..  lower: %f  ..  close: %f  ..  prev_close: %f  ..  trend: %t  ..  value: %f",
	// 	st.ATR.c, basicUpper, st.upper, basicLower, st.lower, candle.Close, st.prev_close, st.IsUptrend(), st.value)

	// upperBand = basicUpperBand < prev upperBand or prev close > prev upperBand ? basicUpperBand : prev upperBand
	if st.upper == 0 || basicUpper < st.upper || st.prev_close > st.upper {
		st.upper = basicUpper
	}

	// lowerBand = basicLowerBand > prev lowerBand or prev close < prev lowerBand ? basicLowerBand : prev lowerBand
	if st.lower == 0 || basicLower > st.lower || st.prev_close < st.lower {
		st.lower = basicLower
	}

	//  hasta acá parece estar todo bien

	// if prev superTrend == prev upperBand
	if st.prev_value == st.prev_upper {
		// 	trendDirection := close > upperBand ? isUpTrend : isDownTrend
		if candle.Close() > st.upper {
			st.value = st.lower
		} else {
			st.value = st.upper
		}
	} else {
		// 	trendDirection := close < lowerBand ? isDownTrend : isUpTrend
		if candle.Close() < st.lower {
			st.value = st.upper
		} else {
			st.value = st.lower
		}
	}

	st.prev_close = candle.Close()
	st.prev_upper = st.upper
	st.prev_value = st.value
}

func (st *SuperTrend) Get() float64 {
	st.mu.Lock()
	defer st.mu.Unlock()

	return st.value
}

// https://es.tradingview.com/support/solutions/43000634738/

// Para calcular las bandas de supertendencia, hay que utilizar las siguientes fórmulas:

// hl2 = (high + low) / 2
// basicUpperBand = hl2 + (multiplier × ATR)
// basicLowerBand = hl2 - (multiplier × ATR)

// upperBand = basicUpperBand < prev upperBand or
//             prev close > prev upperBand ? basicUpperBand : prev upperBand
// lowerBand = basicLowerBand > prev lowerBand or
//             prev close < prev lowerBand ? basicLowerBand : prev lowerBand

// superTrend = trendDirection == isUpTrend ? lowerBand : upperBand
// El parámetro trendDirection se determina en función de que se cumplan las siguientes condiciones:

// Until the ATR value is calculated trendDirection = isDownTrend
// else if prev superTrend == prev upperBand
//     trendDirection := close > upperBand ? isUpTrend : isDownTrend
// else
//     trendDirection := close < lowerBand ? isDownTrend : isUpTrend
