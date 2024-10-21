package tCandle

import (
	"math"

	"github.com/yasseldg/go-simple/types/sFloats"
)

// calculated methods

func (b *Candle) LogReturn() float64 {
	if b.Close() == 0 || b.Open() == 0 {
		return 0
	}
	return sFloats.GetValid(math.Log(b.Close() / b.Open()))
}

func (b *Candle) BodyPerc() float64 {

	total := b.High() - b.Low()
	if total <= 0 {
		return 0
	}

	if b.Open() <= 0 || b.Close() <= 0 {
		return 0
	}

	if b.Close() > b.Open() {
		return ((b.Close() - b.Open()) / total) * 100
	}

	return ((b.Open() - b.Close()) / total) * 100
}
