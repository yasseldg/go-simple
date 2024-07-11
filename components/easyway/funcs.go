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

func GetFunc(name string) ValueFunc {
	switch name {
	case "BuyHead":
		return BuyHead
	case "BuyTail":
		return BuyTail
	case "BuyMiddle":
		return BuyMiddle
	case "SellHead":
		return SellHead
	case "SellTail":
		return SellTail
	case "SellMiddle":
		return SellMiddle
	}
	return nil
}
