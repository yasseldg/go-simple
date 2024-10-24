package indicators

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

func testBBandSuperTrendIter() {

	bb_periods := sInts.NewIter(20, 20, 1)
	bb_deviations := sFloats.NewIter(2.0, 2.0, 0.5, 2)
	bb_config := tIndicator.NewBBConfig(bb_periods, bb_deviations, true)

	st_periods := sInts.NewIter(12, 12, 1)
	st_multipliers := sFloats.NewIter(3.0, 3.0, 0.2, 2)
	st_intervals := tInterval.NewIterLimited()
	st_intervals.Add(tInterval.Interval_D)
	st_at_close := sInts.NewIter(0, 1, 1)
	st_smoothed := sInts.NewIter(0, 1, 1)
	st_config := tIndicator.NewSTConfig(st_periods, st_multipliers,
		st_intervals, st_at_close, st_smoothed, "Test")

	c := 0
	for _, name := range []string{"Sym_1", "Sym_2"} {
		for bb_config.Reset(); bb_config.Next(); {
			for st_config.Reset(); st_config.Next(); {
				c++
				sLog.Info("%s: %5d: BBands: %d -- %f  ..  SuperTrend: %d -- %f", name, c, bb_config.Periods(), bb_config.Deviations(), st_config.Periods(), st_config.Multiplier())
			}
		}
	}
}
