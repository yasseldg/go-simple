package rFilter

// ----- Ts Filters

// Field like Ts, $gte: ts_from  $lt: ts_to
func (f *Filters) TsField(ts_from, ts_to int64, field string) Inter {
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
func (f *Filters) Ts(ts_from, ts_to int64) Inter {
	return f.TsField(ts_from, ts_to, "ts")
}

func (f *Filters) TsIn(tss ...int64) Inter {
	f.Int64_in("ts", tss...)
	return f
}
