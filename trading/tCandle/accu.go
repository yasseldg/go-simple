package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rAccu"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Accu struct {
	rAccu.Inter
}

type mCandle struct {
	rMongo.DefaultModel `bson:",inline"`

	Candle `bson:",inline"`
}

func NewAccu(coll rMongo.InterColl, limit int) (Accu, error) {

	return Accu{
		Inter: rAccu.New(coll, limit),
	}, nil
}

func (iter *Accu) Add(candle *Candle) {
	iter.Inter.Add(&mCandle{Candle: *candle})
}
