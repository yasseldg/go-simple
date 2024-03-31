package sFloats

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
)

// ---- PeriodValues

type PeriodValuesMinMax struct {
	*PeriodValues

	min float64
	max float64
}

func NewPeriodValuesMinMax(period int) *PeriodValuesMinMax {
	return &PeriodValuesMinMax{
		PeriodValues: NewPeriodValues(period),
	}
}

func (pv *PeriodValuesMinMax) Min() float64 {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.min
}

func (pv *PeriodValuesMinMax) Max() float64 {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return pv.max
}

func (pv *PeriodValuesMinMax) Log(qty int) {
	msg := fmt.Sprintf("PeriodValues.Log: periods: %-5d  ..  len: %-5d  ..  min: %f  ..  max: %f", pv.period, len(pv.values), pv.min, pv.max)
	if len(pv.values) < 1 {
		sLog.Warn(msg)
		return
	}
	if len(pv.values) < qty {
		sLog.Warn("%s  ..  values: %v ", msg, pv.values)
		return
	}
	sLog.Warn("PeriodValues.Log: periods: %-5d  ..  len: %-5d  ..  firsts: %v  ..  lasts: %v", pv.period, len(pv.values), pv.values[:3], pv.values[len(pv.values)-3:])
}

func (pv *PeriodValuesMinMax) Add(value float64) {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	if !pv.filled {
		pv.add(value)

		if len(pv.values) == 1 {
			pv.min = value
			pv.max = value
			return
		}

		pv.minMax(value)
		return
	}

	del := pv.values[0]
	pv.values = append(pv.values[1:], value)

	if del == pv.min || del == pv.max {
		pv.min, pv.max = MinMax(pv.values, false)
		return
	}

	pv.minMax(value)
}

func (pv *PeriodValuesMinMax) minMax(value float64) {
	if value > pv.max {
		pv.max = value
		return
	}
	if value < pv.min {
		pv.min = value
	}
}

func (pv *PeriodValuesMinMax) Fill(value float64) bool {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	r := pv.fill(value)

	if len(pv.values) == 1 {
		pv.min = value
		pv.max = value
		return r
	}

	pv.minMax(value)

	return r
}

func (pv *PeriodValuesMinMax) DeepCopy() *PeriodValuesMinMax {
	pv.mu.Lock()
	defer pv.mu.Unlock()

	return &PeriodValuesMinMax{
		PeriodValues: pv.PeriodValues.deepCopy(),
		min:          pv.min,
		max:          pv.max,
	}
}
