package tCandle

import (
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"
)

type OHLC struct {
	Open  float64 `bson:"o" json:"o"`
	High  float64 `bson:"h" json:"h"`
	Low   float64 `bson:"l" json:"l"`
	Close float64 `bson:"c" json:"c"`
}
type OHLCs []OHLC

type OHLCV struct {
	OHLC   `bson:",inline"`
	Volume float64 `bson:"v" json:"v"`
}
type OHLCVs []OHLCV

type Candle struct {
	Ts    int64 `bson:"ts" json:"ts"`
	OHLCV `bson:",inline"`
}
type Candles []Candle

func New(ts int64, open, high, low, close, volume float64) *Candle {
	return &Candle{
		Ts: ts,
		OHLCV: OHLCV{
			OHLC: OHLC{
				Open:  open,
				High:  high,
				Low:   low,
				Close: close},
			Volume: volume},
	}
}

func (c Candle) Log() {
	sLog.Info("Candle: %s  ..  o: %f .. h: %f .. l: %f .. c: %f .. v: %f",
		sTime.ForLog(c.Ts, 0), c.Open, c.High, c.Low, c.Close, c.Volume)
}
