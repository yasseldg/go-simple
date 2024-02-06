package tCandle

import (
	"math"
	"testing"
)

func TestCandle_LogReturn(t *testing.T) {
	tests := []struct {
		name string
		c    Candle
		want float64
	}{
		{"Zero Open and Close", Candle{OHLCV: OHLCV{OHLC: OHLC{Open: 0, Close: 0}}}, 0},
		{"Non-Zero Open and Close", Candle{OHLCV: OHLCV{OHLC: OHLC{Open: 1, Close: 2}}}, 0.693147},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.LogReturn(); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("LogReturn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOHLC_String(t *testing.T) {
	tests := []struct {
		name string
		o    OHLC
		prec int
		want string
	}{
		{"Test OHLC String", OHLC{Open: 1.23, High: 2.34, Low: 0.12, Close: 1.23}, 2, "o: 1.23  ..  h: 2.34  ..  l: 0.12  ..  c: 1.23"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(tt.prec); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOHLCV_String(t *testing.T) {
	tests := []struct {
		name string
		o    OHLCV
		prec int
		want string
	}{
		{"Test OHLCV String", OHLCV{OHLC: OHLC{Open: 1.23, High: 2.34, Low: 0.12, Close: 1.23}, Volume: 123.45}, 2, "o: 1.23  ..  h: 2.34  ..  l: 0.12  ..  c: 1.23  ..  v: 123.45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.String(tt.prec); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCandle_String(t *testing.T) {
	tests := []struct {
		name string
		c    Candle
		prec int
		want string
	}{
		{"Test Candle String", Candle{Ts: 1625140800, OHLCV: OHLCV{OHLC: OHLC{Open: 1.23, High: 2.34, Low: 0.12, Close: 1.23}, Volume: 123.45}}, 2, "Candle: 1625140800 ( 2021.07.01 12:00:00 )  ..  o: 1.23  ..  h: 2.34  ..  l: 0.12  ..  c: 1.23  ..  v: 123.45"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(tt.prec); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
