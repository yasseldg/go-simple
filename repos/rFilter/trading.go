package rFilter

import (
	"github.com/yasseldg/go-simple/trading/tSide"
)

// ----- Trading Filters

func (f *Filters) Sides(sides ...tSide.Side) Inter {
	ints := []int{}
	for _, side := range sides {
		ints = append(ints, int(side))
	}
	f.Int_in("sd", ints...)

	return f
}
