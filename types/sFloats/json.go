package sFloats

import "reflect"

type Json struct {
	From   float64   `json:"from"`
	To     float64   `json:"to"`
	Step   float64   `json:"step"`
	Prec   int       `json:"prec"`
	Values []float64 `json:"values"`
}

func (jn Json) IsZero() bool {
	return reflect.ValueOf(jn).IsZero()
}

func (jn Json) GetIter() *Iter {
	return NewIter(jn.From, jn.To, jn.Step, jn.Prec, jn.Values...)
}
