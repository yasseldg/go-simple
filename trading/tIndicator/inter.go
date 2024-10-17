package tIndicator

import "github.com/yasseldg/go-simple/trading/tIndicator/priceAction"

type InterPriceAction interface {
	priceAction.Inter
}

type InterPriceActionIterConfig interface {
	priceAction.InterIterConfig
}
