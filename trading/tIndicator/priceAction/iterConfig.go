package priceAction

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type IterConfig struct {
	dIter.InterIterConfig

	name     string
	swing    sInts.InterIter
	high_low sInts.InterIter
	closes   sInts.InterIter
}

func NewIterConfig(name string) *IterConfig {
	return &IterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("PriceAction %s", name)),
		name:            name,
	}
}

func (pa *IterConfig) SetSwing(swing sInts.InterIter) InterIterConfig {
	if pa.swing != nil {
		return pa
	}

	pa.swing = swing
	pa.Add(dIter.NewNameConfig("Swing", pa.swing))
	return pa
}

func (pa *IterConfig) SetHighLow(high_low sInts.InterIter) InterIterConfig {
	if pa.high_low != nil {
		return pa
	}

	pa.high_low = high_low
	pa.Add(dIter.NewNameConfig("HighLow", pa.high_low))
	return pa
}

func (pa *IterConfig) SetCloses(closes sInts.InterIter) InterIterConfig {
	if pa.closes != nil {
		return pa
	}

	pa.closes = closes
	pa.Add(dIter.NewNameConfig("Closes", pa.closes))
	return pa
}

func (pa *IterConfig) Swing() bool {
	return pa.swing.Value() == 1
}

func (pa *IterConfig) HighLow() bool {
	return pa.high_low.Value() == 1
}

func (pa *IterConfig) Closes() bool {
	return pa.closes.Value() == 1
}

func (pa *IterConfig) Get() Inter {
	if pa.swing == nil || pa.high_low == nil || pa.closes == nil {
		return nil
	}

	if pa.Closes() && pa.HighLow() {
		return nil
	}

	inter := New().SetSwing(pa.Swing()).SetHighLow(pa.HighLow()).SetCloses(pa.Closes())
	if inter.ConfigNumber() == 0 {
		return nil
	}

	return inter
}

func (pa *IterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(pa.name)
	clone.SetHighLow(pa.high_low.Clone())
	clone.SetSwing(pa.swing.Clone())
	clone.SetCloses(pa.closes.Clone())

	return clone
}
