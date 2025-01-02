package priceAction

import (
	"fmt"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sInts"
)

type iterConfig struct {
	dIter.InterIterConfig

	name     string
	swing    sInts.InterIter
	high_low sInts.InterIter
	closes   sInts.InterIter
}

func NewIterConfig(name string) *iterConfig {
	return &iterConfig{
		InterIterConfig: dIter.NewIterConfig(fmt.Sprintf("PriceAction %s", name)),
		name:            name,
	}
}

func (pa *iterConfig) SetSwing(swing sInts.InterIter) InterIterConfig {
	if pa.swing != nil {
		return pa
	}

	pa.swing = swing
	pa.Add(dIter.NewNameConfig("Swing", pa.swing))
	return pa
}

func (pa *iterConfig) SetHighLow(high_low sInts.InterIter) InterIterConfig {
	if pa.high_low != nil {
		return pa
	}

	pa.high_low = high_low
	pa.Add(dIter.NewNameConfig("HighLow", pa.high_low))
	return pa
}

func (pa *iterConfig) SetCloses(closes sInts.InterIter) InterIterConfig {
	if pa.closes != nil {
		return pa
	}

	pa.closes = closes
	pa.Add(dIter.NewNameConfig("Closes", pa.closes))
	return pa
}

func (pa *iterConfig) Swing() bool {
	return pa.swing.Value() == 1
}

func (pa *iterConfig) HighLow() bool {
	return pa.high_low.Value() == 1
}

func (pa *iterConfig) Closes() bool {
	return pa.closes.Value() == 1
}

func (pa *iterConfig) Get() (Inter, error) {
	if pa.swing == nil {
		return nil, fmt.Errorf("swing is required")
	}

	if pa.high_low == nil {
		return nil, fmt.Errorf("high_low is required")
	}

	if pa.closes == nil {
		return nil, fmt.Errorf("closes is required")
	}

	if pa.Closes() && pa.HighLow() {
		return nil, fmt.Errorf("closes and high_low")
	}

	inter := New().SetSwing(pa.Swing()).SetHighLow(pa.HighLow()).SetCloses(pa.Closes())
	if inter.ConfigNumber() == 0 {
		return nil, fmt.Errorf("invalid config")
	}

	return inter, nil
}

func (pa *iterConfig) Clone() InterIterConfig {

	clone := NewIterConfig(pa.name)
	clone.SetHighLow(pa.high_low.Clone())
	clone.SetSwing(pa.swing.Clone())
	clone.SetCloses(pa.closes.Clone())

	return clone
}
