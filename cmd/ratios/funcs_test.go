package ratios

import (
	"testing"

	"github.com/yasseldg/go-simple/trading/tFunc"
	"github.com/yasseldg/go-simple/trading/tSide"
)

func TestPrueba(t *testing.T) {
	high := 10245.0
	entry_price := 10069.0
	low := 9952.0

	side := tSide.Buy
	tp_price := high
	sl_price := low

	ratio := tFunc.Ratio(entry_price, sl_price, tp_price)
	expectedRatio := (tp_price / sl_price)

	if ratio != expectedRatio {
		t.Errorf("Expected ratio %f, got %f", expectedRatio, ratio)
	}

	entry := tFunc.EntryByRatio(ratio, sl_price, tp_price)
	expectedEntry := entry_price

	if entry != expectedEntry {
		t.Errorf("Expected entry %f, got %f", expectedEntry, entry)
	}

	take_profit := tFunc.TakeProfitByRatio(entry_price, sl_price, ratio)
	expectedTakeProfit := tp_price

	if take_profit != expectedTakeProfit {
		t.Errorf("Expected take profit %f, got %f", expectedTakeProfit, take_profit)
	}

	stop_loss := tFunc.StopLossByRatio(entry_price, tp_price, ratio)
	expectedStopLoss := sl_price

	if stop_loss != expectedStopLoss {
		t.Errorf("Expected stop loss %f, got %f", expectedStopLoss, stop_loss)
	}

	side = tSide.Sell
	sl_price = high
	tp_price = low

	ratio = tFunc.Ratio(entry_price, sl_price, tp_price)
	expectedRatio = (tp_price / sl_price)

	if ratio != expectedRatio {
		t.Errorf("Expected ratio %f, got %f", expectedRatio, ratio)
	}

	entry = tFunc.EntryByRatio(ratio, sl_price, tp_price)
	expectedEntry = entry_price

	if entry != expectedEntry {
		t.Errorf("Expected entry %f, got %f", expectedEntry, entry)
	}

	take_profit = tFunc.TakeProfitByRatio(entry_price, sl_price, ratio)
	expectedTakeProfit = tp_price

	if take_profit != expectedTakeProfit {
		t.Errorf("Expected take profit %f, got %f", expectedTakeProfit, take_profit)
	}

	stop_loss = tFunc.StopLossByRatio(entry_price, tp_price, ratio)
	expectedStopLoss = sl_price

	if stop_loss != expectedStopLoss {
		t.Errorf("Expected stop loss %f, got %f", expectedStopLoss, stop_loss)
	}

	risk := 0.784
	reward := 0.149
	expectedRatio = tFunc.RatioFromRiskReward(risk, reward)

	if tFunc.RatioFromRiskReward(risk, reward) != expectedRatio {
		t.Errorf("Expected ratio %f, got %f", expectedRatio, tFunc.RatioFromRiskReward(risk, reward))
	}

	expectedRatio = tFunc.RatioFromRiskReward(reward, risk)

	if tFunc.RatioFromRiskReward(reward, risk) != expectedRatio {
		t.Errorf("Expected ratio %f, got %f", expectedRatio, tFunc.RatioFromRiskReward(reward, risk))
	}
}
