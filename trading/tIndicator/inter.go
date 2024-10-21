package tIndicator

import (
	"github.com/yasseldg/go-simple/trading/tIndicator/priceAction"
)

// Price Action
type InterPriceAction interface {
	priceAction.Inter
}

type InterPriceActionIterConfig interface {
	priceAction.InterIterConfig
}

// Super Trend
// type InterSuperTrend interface {
// 	superTrend.Inter
// }

type InterSuperTrendIterConfig interface {
	InterSTConfig
}
