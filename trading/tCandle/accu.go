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

func (accu *Accu) Upsert(candle *Candle) error {

	filter := NewFilters()
	filter.Ts(candle.Ts(), candle.Ts())

	return accu.Inter.Coll().Filters(filter).Upsert(
		&model{Candle: *candle},
	)
}
