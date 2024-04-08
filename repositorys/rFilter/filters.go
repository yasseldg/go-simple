package rFilter

import (
	"github.com/yasseldg/go-simple/trading/tSide"
)

type Filters struct {
	Inter
}

func New(inter Inter) *Filters {
	return &Filters{Inter: inter}
}

func (f *Filters) Clone() *Filters {
	return &Filters{Inter: f.Inter.Clone()}
}

// ----- Ts Filters

// Field like Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) TsField(ts_from, ts_to int64, field string) *Filters {
	if ts_from > 0 {
		if ts_to > 0 {
			// sLog.Debug("TsField: ts_from: %d  ts_to: %d", ts_from, ts_to)
			return f.Int64_gte_lt(field, ts_from, ts_to)
		}
		// sLog.Debug("TsField: ts_from: %d", ts_from)
		return f.Int64_gte(field, ts_from)
	}
	if ts_to > 0 {
		// sLog.Debug("TsField: ts_to: %d", ts_to)
		return f.Int64_lt(field, ts_to)
	}
	return f
}

// Set Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) Ts(ts_from, ts_to int64) *Filters {
	return f.TsField(ts_from, ts_to, "ts")
}

func (f *Filters) TsIn(tss ...int64) *Filters {
	f.Int64_in("ts", tss...)
	return f
}

// ----- States Filters

func (f *Filters) States(states ...string) *Filters { f.String_in("st", states...); return f }

func (f *Filters) NotStates(states ...string) *Filters { f.String_nin("st", states...); return f }

// ----- Trading Filters

func (f *Filters) Sides(sides ...tSide.Side) *Filters {
	ints := []int{}
	for _, side := range sides {
		ints = append(ints, int(side))
	}
	f.Int_in("sd", ints...)

	return f
}
