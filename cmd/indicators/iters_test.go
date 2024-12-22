package indicators

import (
	"testing"

	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

func TestBBandSuperTrendIter(t *testing.T) {
	bb_config := tIndicator.NewBBIterConfig("Test")
	bb_config.SetPeriods(sInts.NewIter(20, 20, 1))
	bb_config.SetDeviations(sFloats.NewIter(2.0, 2.0, 0.5, 2))

	st_config := tIndicator.NewSuperTrendIterConfig("Test")
	st_config.SetPeriods(sInts.NewIter(12, 12, 1))
	st_config.SetMultiplier(sFloats.NewIter(3.0, 3.0, 0.2, 2))
	st_config.SetSmoothed(sInts.NewIter(0, 1, 1))

	c := 0
	for _, name := range []string{"Sym_1", "Sym_2"} {
		for bb_config.Reset(); bb_config.Next(); {
			for st_config.Reset(); st_config.Next(); {
				c++
				st_config.Log(name)
			}
		}
	}

	if c != 8 {
		t.Errorf("Expected 8 iterations, got %d", c)
	}
}
