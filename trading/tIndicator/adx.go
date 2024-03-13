package tIndicator

import (
	"math"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sTime"
)

// El Índice de Movimiento Direccional Promedio (ADX, por sus siglas en inglés)

type ADX struct {
	ATR

	plusDMs  sFloats.SmoothedAverage
	minusDMs sFloats.SmoothedAverage

	value    sFloats.SmoothedAverage
	historic sFloats.Average
}

func NewADX(period int) *ADX {
	return &ADX{
		ATR: *NewATR(period),

		plusDMs:  *sFloats.NewSmoothedAverage(period),
		minusDMs: *sFloats.NewSmoothedAverage(period),

		value:    *sFloats.NewSmoothedAverage(period),
		historic: *sFloats.NewAverage(),
	}
}

func (adx *ADX) Log() {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	sLog.Info("ADX %d: %s  ..  TRs: %f  ..  pDMs: %f  ..  mDMs: %f  .. adx: %f  ..  %s", adx.c, sTime.ForLog(adx.prev.Ts, 0), adx.trs.Value(), adx.plusDMs.Value(), adx.minusDMs.Value(), adx.value.Value(), adx.historic.String())
}

func (adx *ADX) Add(candle tCandle.Candle) {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	if adx.prev.Close > 0 {
		adx.calcDMs(candle)

		adx.ATR.add(candle)

		adx.calcDIs()
	}

	adx.prev = candle
}

func (adx *ADX) Value() float64 {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	return adx.value.Value()
}

func (adx *ADX) Historic() float64 {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	return adx.historic.Calc()
}

func (adx *ADX) calcDMs(candle tCandle.Candle) {

	plusDM := candle.High - adx.prev.High
	minusDM := adx.prev.Low - candle.Low

	if plusDM <= minusDM || plusDM <= 0 {
		plusDM = 0
	}
	if minusDM <= plusDM || minusDM <= 0 {
		minusDM = 0
	}

	adx.plusDMs.AddPos(plusDM)
	adx.minusDMs.AddPos(minusDM)
}

func (adx *ADX) calcDIs() {

	if !adx.trs.Filled() || !adx.plusDMs.Filled() || !adx.minusDMs.Filled() {
		return
	}

	plusDi := (adx.plusDMs.Value() / adx.trs.Value()) * 100
	minusDi := (adx.minusDMs.Value() / adx.trs.Value()) * 100

	dx := (math.Abs(plusDi-minusDi) / math.Abs(plusDi+minusDi)) * 100

	adx.value.AddPos(dx)

	// sLog.Debug("+Di: %f  .. -Di %f  .. Dx: %f", plusDi, minusDi, dx)

	if adx.value.Filled() {
		adx.historic.Add(adx.value.Value())
	}
}
