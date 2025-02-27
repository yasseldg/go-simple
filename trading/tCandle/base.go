package tCandle

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"
	"github.com/yasseldg/go-simple/types/sTime"
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

func New(ts int64, open, high, low, close, volume float64) *Candle {
	return &Candle{
		M_ts: ts,
		OHLCV: OHLCV{
			OHLC: OHLC{
				M_open:  open,
				M_high:  high,
				M_low:   low,
				M_close: close},
			M_volume: volume},
	}
}

func (b *OHLC) String(prec int) string {
	return fmt.Sprintf("o: %.*f  ..  h: %.*f  ..  l: %.*f  ..  c: %.*f", prec, b.M_open, prec, b.M_high, prec, b.M_low, prec, b.M_close)
}

func (b *OHLCV) String(prec int) string {
	return fmt.Sprintf("%s  ..  v: %.*f", b.OHLC.String(prec), prec, b.M_volume)
}

func (b *Candle) String(prec int) string {
	return fmt.Sprintf("%s  ..  %s", sTime.ForLog(b.M_ts, 0), b.OHLCV.String(prec))
}

func (b *Candle) Log(prec int) {
	sLog.Info("Candle: %s", b.String(prec))
}

func (b *Candle) Ts() int64 {
	return b.M_ts
}

func (b *Candle) Open() float64 {
	return b.M_open
}

func (b *Candle) High() float64 {
	return b.M_high
}

func (b *Candle) Low() float64 {
	return b.M_low
}

func (b *Candle) Close() float64 {
	return b.M_close
}

func (b *Candle) Volume() float64 {
	return b.M_volume
}

func (b *Candle) GetModel() *Candle {
	return b
}

func (b *Candle) Fill(ts int64, repo rMongo.InterRepo) error {

	var candle model
	err := repo.Clone().First(ts, (ts + 1), &candle)
	if err != nil {
		return fmt.Errorf("coll.First( %d ): %s", ts, err)
	}

	*b = candle.Candle

	return nil
}
