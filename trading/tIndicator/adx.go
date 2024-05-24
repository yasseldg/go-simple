package tIndicator

import (
	"fmt"
	"math"
	"sync"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tCandle"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type InterADX interface {
	InterATR
}

// El Índice de Movimiento Direccional Promedio (ADX, por sus siglas en inglés)

type ADX struct {
	InterATR

	mu sync.Mutex

	plusDMs  sFloats.SmoothedAverage
	minusDMs sFloats.SmoothedAverage

	value    sFloats.SmoothedAverage
	historic sFloats.Average
}

func NewADX(period int) *ADX {
	return &ADX{
		InterATR: NewSmATR(period),

		plusDMs:  *sFloats.NewSmoothedAverage(period),
		minusDMs: *sFloats.NewSmoothedAverage(period),

		value:    *sFloats.NewSmoothedAverage(period),
		historic: *sFloats.NewAverage(),
	}
}

func (adx *ADX) String() string {
	return fmt.Sprintf("%s  ..  pDMs: %f  ..  mDMs: %f  ..  adx: %f  ..  %s",
		adx.InterATR.String(), adx.plusDMs.Value(), adx.minusDMs.Value(), adx.value.Value(), adx.historic.String())
}

func (adx *ADX) Log() {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	sLog.Info("ADX: %s", adx.String())
}

func (adx *ADX) Add(candle tCandle.Inter) {
	adx.mu.Lock()
	defer adx.mu.Unlock()

	if adx.Prev().Close() <= 0 {
		adx.InterATR.Add(candle)
		return
	}

	adx.calcDMs(candle)

	adx.InterATR.Add(candle)

	adx.calcDIs()
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

func (adx *ADX) calcDMs(candle tCandle.Inter) {

	plusDM := candle.High() - adx.Prev().High()
	minusDM := adx.Prev().Low() - candle.Low()

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

	if !adx.InterATR.Filled() || !adx.plusDMs.Filled() || !adx.minusDMs.Filled() {
		return
	}

	plusDi := (adx.plusDMs.Value() / adx.InterATR.Get()) * 100
	minusDi := (adx.minusDMs.Value() / adx.InterATR.Get()) * 100

	dx := (math.Abs(plusDi-minusDi) / math.Abs(plusDi+minusDi)) * 100

	adx.value.AddPos(dx)

	// sLog.Debug("+Di: %f  .. -Di %f  .. Dx: %f", plusDi, minusDi, dx)

	if adx.value.Filled() {
		adx.historic.Add(adx.value.Value())
	}
}
