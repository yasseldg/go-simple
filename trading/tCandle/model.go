package tCandle

import (
	"github.com/yasseldg/go-simple/repos/rFilter"
	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo"
)

type OHLC struct {
	M_open  float64 `bson:"o" json:"o"`
	M_high  float64 `bson:"h" json:"h"`
	M_low   float64 `bson:"l" json:"l"`
	M_close float64 `bson:"c" json:"c"`
}

type OHLCV struct {
	OHLC     `bson:",inline"`
	M_volume float64 `bson:"v" json:"v"`
}

type Candle struct {
	OHLCV `bson:",inline"`
	M_ts  int64 `bson:"ts" json:"ts"`
}
type Candles []*Candle

type model struct {
	rMongo.Model `bson:",inline"`

	Candle `bson:",inline"`
}

// filters

type Filters struct{ rFilter.Filters }

func NewFilters() *Filters {
	return &Filters{Filters: *rMongo.NewFilter()}
}

func (f *Filters) High_gt(high float64) *Filters { f.Gt("h", high); return f }

func (f *Filters) Low_lt(low float64) *Filters { f.Lt("l", low); return f }

// indexes

func Indexes() rIndex.Indexes {
	return rIndex.Indexes{
		rIndex.New(rMongo.NewSort().TsAsc(), true),
	}
}
