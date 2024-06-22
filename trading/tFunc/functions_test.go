package tFunc

import (
	"testing"

	"github.com/yasseldg/go-simple/trading/tSide"
)

func TestTakeProfitStoploss(t *testing.T) {
	testCases := []struct {
		name             string
		entry_price      float64
		take_profit_perc float64
		stop_loss_perc   float64
		side             tSide.Side
		expected_tp      float64
		expected_sl      float64
	}{
		{
			name:             "Buy Side Test",
			entry_price:      100.0,
			take_profit_perc: 10.0,
			stop_loss_perc:   5.0,
			side:             tSide.Buy,
			expected_tp:      110.0,
			expected_sl:      95.0,
		},
		{
			name:             "Sell Side Test",
			entry_price:      100.0,
			take_profit_perc: 10.0,
			stop_loss_perc:   5.0,
			side:             tSide.Sell,
			expected_tp:      90.0,
			expected_sl:      105.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tp, sl := TakeProfitStopLoss(tc.entry_price, tc.take_profit_perc, tc.stop_loss_perc, tc.side)
			if tp != tc.expected_tp || sl != tc.expected_sl {
				t.Errorf("TakeProfitStoploss() = %v, %v, want %v, %v", tp, sl, tc.expected_tp, tc.expected_sl)
			}
		})
	}
}
