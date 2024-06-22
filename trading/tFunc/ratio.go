package tFunc

import (
	"math"

	"github.com/yasseldg/go-simple/types/sFloats"
)

func RatioFromRiskReward(risk, reward float64) float64 {
	return sFloats.GetValid(reward / risk)
}

func Ratio(entry_price, sl_price, tp_price float64) float64 {

	return sFloats.GetValid(math.Abs(tp_price-entry_price) / math.Abs(entry_price-sl_price))
}

func EntryByRatio(ratio, sl_price, tp_price float64) float64 {

	return (tp_price + (sl_price * ratio)) / (1 + ratio)
}

func StopLossByRatio(entry_price, tp_price, ratio float64) float64 {

	return entry_price - ((tp_price - entry_price) / ratio)
}

func TakeProfitByRatio(entry_price, sl_price, ratio float64) float64 {

	return entry_price + (ratio * (entry_price - sl_price))
}
