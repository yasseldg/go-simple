package superTrend

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/trading/tIndicator/atr"
	"github.com/yasseldg/go-simple/types/sStrings"
	"github.com/yasseldg/go-simple/types/sTime"
)

type base struct {
	atr.Inter

	mu sync.Mutex

	multiplier float64

	prev_close float64
	prev_upper float64
	prev_value float64

	upper float64
	lower float64

	value float64

	ts_last int64
}

func New(periods int, multiplier float64, smoothed bool) *base {
	supertrend := &base{
		multiplier: multiplier,
	}

	if smoothed {
		supertrend.Inter = atr.NewSmoothed(periods)
	} else {
		supertrend.Inter = atr.NewAvg(periods)
	}

	return supertrend
}

func (st *base) Config() string {
	return fmt.Sprintf("period: %d  ..  multiplier: %.2f", st.Periods(), st.multiplier)
}

func (st *base) String() string {
	v := ""
	if st.IsUptrend() {
		v = sStrings.Colored(sStrings.Green, fmt.Sprintf("%f", st.value))
	} else {
		v = sStrings.Colored(sStrings.Red, fmt.Sprintf("%f", st.value))
	}

	return fmt.Sprintf("SuperTrend %d: period: %d  ..  multiplier: %.2f  ..  %s  ..  %s",
		st.Inter.Count(), st.Periods(), st.multiplier, sTime.ForLog(st.Inter.Prev().Ts(), 0), v)
}

func (st *base) Log() {
	st.mu.Lock()
	defer st.mu.Unlock()

	sLog.Info(st.String())
}

func (st *base) Multiplier() float64 {
	return st.multiplier
}

func (st *base) IsUptrend() bool {
	return st.value == st.lower
}

func (st *base) IsDowntrend() bool {
	return st.value == st.upper
}

func (st *base) Add(candle tCandle.Inter) {
	st.mu.Lock()
	defer st.mu.Unlock()

	if st.ts_last >= candle.Ts() {
		return
	}

	st.Inter.Add(candle)

	if st.Inter.Get() == 0 {
		return
	}

	hl2 := (candle.High() + candle.Low()) / 2

	multAtr := st.multiplier * st.Inter.Get()

	basicUpper := hl2 + multAtr
	basicLower := hl2 - multAtr

	if st.upper == 0 || basicUpper < st.upper || st.prev_close > st.upper {
		st.upper = basicUpper
	}

	if st.lower == 0 || basicLower > st.lower || st.prev_close < st.lower {
		st.lower = basicLower
	}

	if st.prev_value == st.prev_upper {
		if candle.Close() > st.upper {
			st.value = st.lower
		} else {
			st.value = st.upper
		}
	} else {
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

func (st *base) Get() float64 {
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
