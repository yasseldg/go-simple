package indicators

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

func testBBandSuperTrendIter() {

	bb_periods := sInts.NewIter(18, 20, 1)
	bb_deviations := sFloats.NewIter(1.0, 2.0, 0.5)
	bb_config := tIndicator.NewBBConfig(bb_periods, bb_deviations, true)

	st_periods := sInts.NewIter(12, 15, 1)
	st_multipliers := sFloats.NewIter(2.0, 3.0, 0.2)
	st_config := tIndicator.NewSTConfig(st_periods, st_multipliers, true, false)

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
