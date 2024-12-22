package iters

import (
	"testing"

	"github.com/yasseldg/go-simple/data/dIter"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
	"github.com/yasseldg/go-simple/trading/tIndicator"
	"github.com/yasseldg/go-simple/trading/tInterval"
	"github.com/yasseldg/go-simple/trading/tSymbol"
)

func TestConfigs(t *testing.T) {
	st_config := configST()

	if st_config == nil {
		t.Error("Expected non-nil st_config")
	}

	floats := configFloats()
	if floats == nil {
		t.Error("Expected non-nil floats")
	}

	ints := configInts()
	if ints == nil {
		t.Error("Expected non-nil ints")
	}

	symbs := configSymbols()
	if symbs == nil {
		t.Error("Expected non-nil symbs")
	}

	intervals := configIntervals()
	if intervals == nil {
		t.Error("Expected non-nil intervals")
	}

	floats_0 := sFloats.NewIter(0, 0, 0.2, 2)
	if floats_0 == nil {
		t.Error("Expected non-nil floats_0")
	}

	ints_1 := sInts.NewIter(9, 8, 1)
	if ints_1 == nil {
		t.Error("Expected non-nil ints_1")
	}

	ints_2 := sInts.NewIter(0, 0, 1)
	if ints_2 == nil {
		t.Error("Expected non-nil ints_2")
	}

	config := dIter.NewIterConfig("Basics")
	if config == nil {
		t.Error("Expected non-nil config")
	}

	config.Add(dIter.NewNameConfig("Symbols", symbs))
	config.Add(dIter.NewNameConfig("ST Entry", st_config))
	config.Add(dIter.NewNameConfig("Ints 1", ints_1))
	config.Add(dIter.NewNameConfig("Ints 2", ints_2))
	config.Add(dIter.NewNameConfig("Intervals", intervals))
	config.Add(dIter.NewNameConfig("Floats 0", floats_0))

	for config.Reset(); config.Next(); {
		// Just iterating to ensure no panics
	}
}

func TestConfigST(t *testing.T) {
	st_config := configST()

	if st_config == nil {
		t.Error("Expected non-nil st_config")
	}

	for st_config.Reset(); st_config.Next(); {
		// Just iterating to ensure no panics
	}
}

func TestConfigFloats(t *testing.T) {
	floats := configFloats()

	if floats == nil {
		t.Error("Expected non-nil floats")
	}

	for floats.Reset(); floats.Next(); {
		// Just iterating to ensure no panics
	}
}

func TestConfigInts(t *testing.T) {
	ints := configInts()

	if ints == nil {
		t.Error("Expected non-nil ints")
	}

	for ints.Reset(); ints.Next(); {
		// Just iterating to ensure no panics
	}
}

func TestConfigSymbols(t *testing.T) {
	symbs := configSymbols()

	if symbs == nil {
		t.Error("Expected non-nil symbs")
	}

	for symbs.Reset(); symbs.Next(); {
		// Just iterating to ensure no panics
	}
}

func TestConfigIntervals(t *testing.T) {
	intervals := configIntervals()

	if intervals == nil {
		t.Error("Expected non-nil intervals")
	}

	for intervals.Reset(); intervals.Next(); {
		// Just iterating to ensure no panics
	}
}
