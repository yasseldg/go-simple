package tFunc

import (
	"math"
	"testing"

	"github.com/yasseldg/go-simple/trading/tSide"
)

const tolerance = 1e-9

func floatEquals(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestTakeProfit(t *testing.T) {
	testCases := []struct {
		name        string
		entry_price float64
		tp_perc     float64
		side        tSide.Side
		expected    float64
	}{
		{
			name:        "Buy Side Take Profit",
			entry_price: 100.0,
			tp_perc:     10.0,
			side:        tSide.Buy,
			expected:    110.0,
		},
		{
			name:        "Sell Side Take Profit",
			entry_price: 100.0,
			tp_perc:     10.0,
			side:        tSide.Sell,
			expected:    90.0,
		},
	}

	println("tolerance: %f", tolerance)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := TakeProfit(tc.entry_price, tc.tp_perc, tc.side)
			if !floatEquals(actual, tc.expected, tolerance) {
				t.Errorf("TakeProfit() = %v, want %v", actual, tc.expected)
			}
		})
	}
}

func TestStopLoss(t *testing.T) {
	testCases := []struct {
		name        string
		entry_price float64
		sl_perc     float64
		side        tSide.Side
		expected    float64
	}{
		{
			name:        "Buy Side Stop Loss",
			entry_price: 100.0,
			sl_perc:     5.0,
			side:        tSide.Buy,
			expected:    95.0,
		},
		{
			name:        "Sell Side Stop Loss",
			entry_price: 100.0,
			sl_perc:     5.0,
			side:        tSide.Sell,
			expected:    105.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := StopLoss(tc.entry_price, tc.sl_perc, tc.side)
			if !floatEquals(actual, tc.expected, tolerance) {
				t.Errorf("StopLoss() = %v, want %v", actual, tc.expected)
			}
		})
	}
}
