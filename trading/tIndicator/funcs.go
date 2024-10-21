package tIndicator

import (
	"github.com/yasseldg/go-simple/trading/tIndicator/priceAction"
)

func NewPriceAction() priceAction.Inter {
	return priceAction.New()
}

func NewPriceActionIterConfig(name string) priceAction.InterIterConfig {
	return priceAction.NewIterConfig(name)
}
