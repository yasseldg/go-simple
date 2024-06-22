package easyway

type ValueFunc func(InterValueFuncs) float64

func BuyHead(ew InterValueFuncs) float64 {
	return ew.BuyHead()
}

func BuyTail(ew InterValueFuncs) float64 {
	return ew.BuyTail()
}

func BuyMiddle(ew InterValueFuncs) float64 {
	return ew.BuyMiddle()
}

func SellHead(ew InterValueFuncs) float64 {
	return ew.SellHead()
}

func SellTail(ew InterValueFuncs) float64 {
	return ew.SellTail()
}

func SellMiddle(ew InterValueFuncs) float64 {
	return ew.SellMiddle()
}
