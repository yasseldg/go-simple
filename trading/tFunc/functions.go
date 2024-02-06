package tFunc

import (
	"github.com/yasseldg/go-simple/trading/tSide"
	"github.com/yasseldg/go-simple/types/sFloats"
)

// TakeProfitStoploss: returns take_profit and stop_loss prices
func TakeProfitStoploss(entry_price, take_profit_perc, stop_loss_perc float64, side tSide.Side) (take_profit, stop_loss float64) {
	switch side {
	case tSide.Buy:
		take_profit = sFloats.GetWithPercent(entry_price, take_profit_perc, true)
		stop_loss = sFloats.GetWithPercent(entry_price, stop_loss_perc, false)

	case tSide.Sell:
		take_profit = sFloats.GetWithPercent(entry_price, take_profit_perc, false)
		stop_loss = sFloats.GetWithPercent(entry_price, stop_loss_perc, true)
	}
	return
}
