package iters

import (
	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/trading/tSymbol"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

func Configs() {

	st_config := configST()

	// return

	println()
	floats := configFloats()
	floats.Log("Floats Out")

	println()
	ints := configInts()
	ints.Log("Ints Out")

	println()
	symbs := configSymbols()
	symbs.Log("Symbols Out")

	println()
	intervals := configIntervals()
	intervals.Log("Intervals Out")

	println()
	floats_0 := sFloats.NewIter(0, 0, 0.2, 2)
	floats_0.Log("Floats_0")

	ints_1 := sInts.NewIter(9, 8, 1)
	ints_1.Log("Ints_1")

	ints_2 := sInts.NewIter(0, 0, 1)
	ints_2.Log("Ints_2")

	println()

	config := dIter.NewIterConfig("Basics")

	println()
	config.Add(dIter.NewNameConfig("Symbols", symbs))
	config.Log("")
	println()
	config.Add(dIter.NewNameConfig("ST Entry", st_config))
	config.Log("")
	println()
	config.Add(dIter.NewNameConfig("Ints 1", ints_1))
	config.Log("")
	println()
	config.Add(dIter.NewNameConfig("Ints 2", ints_2))
	config.Log("")
	// println()
	// config.Add(dIter.NewNameConfig("Floats SL", floats))
	// config.Log("")
	// println()
	// config.Add(dIter.NewNameConfig("Ints Periods", ints))
	// config.Log("")
	println()
	config.Add(dIter.NewNameConfig("Intervals", intervals))
	config.Log("")
	println()
	config.Add(dIter.NewNameConfig("Floats 0", floats_0))
	config.Log("")

	println()
	sLog.Warn("Reset Config")
	config.Reset()
	config.Log("")
	println()
	println()

	// return

	for config.Reset(); config.Next(); {
		println()
		config.Log("")
		println()
	}
}

func configST() dIter.InterConfig {
	st_periods := sInts.NewIter(11, 12, 1)
	st_multipliers := sFloats.NewIter(2.8, 3.0, 0.2, 2)
	st_config := tIndicator.NewSTConfig(st_periods, st_multipliers, true, false, "Test")

	for st_config.Reset(); st_config.Next(); {
		st_config.Log("")
	}

	return st_config
}

func configFloats() dIter.InterConfig {
	floats := sFloats.NewIter(1.5, 1.7, 0.2, 2)
	floats.Log("Floats 1")

	println()
	run(floats, "floats run 1")

	floats.Reset()
	floats.Log("Floats Reset")

	println()
	run(floats, "floats run 2")

	return floats
}

func configInts() dIter.InterConfig {
	ints := sInts.NewIter(2, 4, 1, 6, 8)
	ints.Log("Ints 1")

	println()
	run(ints, "ints run 1")

	ints.Reset()
	ints.Log("Ints Reset")

	println()
	run(ints, "ints run 2")

	return ints
}

func configSymbols() dIter.InterConfig {
	symbs := tSymbol.NewIterLimited()
	symbs.Log("Symbols 1 empty")

	symbs.Add(tSymbol.New("Bybit", "BTCUSD"))
	// symbs.Add(tSymbol.New("Bybit", "LTCUSD"))
	// symbs.Add(tSymbol.New("Bybit", "ETHUSD"))
	// symbs.Add(tSymbol.New("Bybit", "XRPUSD"))
	// symbs.Add(tSymbol.New("Bybit", "EOSUSD"))
	symbs.Log("Symbols 1")

	println()
	runSymbols(symbs, "symbols run 1")

	symbs.Reset()
	symbs.Log("Symbols Reset")

	println()
	runSymbols(symbs, "symbols run 2")

	return symbs
}

func configIntervals() dIter.InterConfig {
	intervals := tInterval.NewIterLimited()
	intervals.Log("Intervals 1 empty")

	// intervals.Add(tInterval.Interval_1m)
	intervals.Add(tInterval.Interval_15m)
	intervals.Add(tInterval.Interval_1h)

	intervals.Log("Intervals 1")

	println()
	runIntervals(intervals, "intervals run 1")

	intervals.Reset()
	intervals.Log("Intervals Reset")

	println()
	runIntervals(intervals, "intervals run 2")

	return intervals
}
