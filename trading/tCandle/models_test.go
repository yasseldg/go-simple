package tCandle

import (
	"testing"
)

func TestOHLC(t *testing.T) {
	ohlc := OHLC{Open: 1.0, High: 2.0, Low: 0.5, Close: 1.5}

	if ohlc.Open != 1.0 {
		t.Errorf("Expected Open to be 1.0, got %f", ohlc.Open)
	}

	if ohlc.High != 2.0 {
		t.Errorf("Expected High to be 2.0, got %f", ohlc.High)
	}

	if ohlc.Low != 0.5 {
		t.Errorf("Expected Low to be 0.5, got %f", ohlc.Low)
	}

	if ohlc.Close != 1.5 {
		t.Errorf("Expected Close to be 1.5, got %f", ohlc.Close)
	}
}

func TestOHLCV(t *testing.T) {
	ohlcv := OHLCV{OHLC: OHLC{Open: 1.0, High: 2.0, Low: 0.5, Close: 1.5}, Volume: 1000.0}

	if ohlcv.Volume != 1000.0 {
		t.Errorf("Expected Volume to be 1000.0, got %f", ohlcv.Volume)
	}
}

func TestCandle(t *testing.T) {
	candle := Candle{Ts: 123456789, OHLCV: OHLCV{OHLC: OHLC{Open: 1.0, High: 2.0, Low: 0.5, Close: 1.5}, Volume: 1000.0}}

	if candle.Ts != 123456789 {
		t.Errorf("Expected Ts to be 123456789, got %d", candle.Ts)
	}
}
