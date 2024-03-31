package tCandle

import (
	"github.com/yasseldg/go-simple/repositorys/rAccu"
	"github.com/yasseldg/go-simple/repositorys/rMongo"

	"github.com/yasseldg/mgm/v4"
)

type Accu struct {
	rAccu.Accu
}

type mgmCandle struct {
	mgm.DefaultModel `bson:",inline"`

	Candle `bson:",inline"`
}

func NewAccu(coll rMongo.Collection, limit int) (Accu, error) {

	return Accu{Accu: rAccu.New(coll, limit)}, nil
}

func (iter *Accu) Add(candle *Candle) {
	iter.Accu.Add(&mgmCandle{Candle: *candle})
}
