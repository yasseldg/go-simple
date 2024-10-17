package sFloats

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sInts"
)

type InterIter interface {
	dIter.Inter

	Value() float64
	Count() int
	Reset()
}

type Iter struct {
	sInts.InterIter

	from   float64
	to     float64
	step   float64
	prec   int
	values []float64

	current float64
}

func NewIter(from, to, step float64, prec int, values ...float64) *Iter {

	int_values := make([]int64, len(values))
	for i, v := range values {
		int_values[i] = sInts.InflateFloat64(v, prec)
	}

	return &Iter{
		InterIter: sInts.NewIter(sInts.InflateFloat64(from, prec), sInts.InflateFloat64(to, prec),
			sInts.InflateFloat64(step, prec), int_values...),

		from:   from,
		to:     to,
		step:   step,
		prec:   prec,
		values: values,
	}
}

func (b *Iter) String(name string) string {
	return fmt.Sprintf("%s ..  from: %f  ..  to: %f  ..  step: %f  ..  values: %v  ..  value: %f",
		b.InterIter.String(name), b.from, b.to, b.step, b.values, b.current)
}

func (b *Iter) Log(name string) {
	sLog.Info(b.String(name))
}

func (b *Iter) Value() float64 {
	return b.current
}

func (b *Iter) Next() bool {

	if b.InterIter.Next() {
		b.current = sInts.DeflateFloat64(b.InterIter.Value(), b.prec)
		return true
	}

	return false
}

func (b *Iter) Reset() {
	b.current = 0
	b.InterIter.Reset()
}

func (b *Iter) Clone() InterIter {
	return NewIter(b.from, b.to, b.step, b.prec, b.values...)
}
