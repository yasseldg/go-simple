package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rAccu"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type Accu struct {
	rAccu.Inter
}

func NewAccu(coll rMongo.InterRepo, limit int) (*Accu, error) {
	return &Accu{
		Inter: rAccu.New(coll, limit),
	}, nil
}

func (accu *Accu) AddCandle(candle *Candle) {
	accu.Inter.Add(&model{Candle: *candle})
}
