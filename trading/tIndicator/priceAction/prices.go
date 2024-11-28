package priceAction

func (pa *base) UpPrice() float64 {

	if pa.new_low == 0 {
		return 0
	}

	if pa.state == uptrend || pa.state == downtrend {
		return 0
	}

	return pa.last_high
}

func (pa *base) DownPrice() float64 {

	if pa.new_high == 0 {
		return 0
	}

	if pa.state == uptrend || pa.state == downtrend {
		return 0
	}

	return pa.last_low
}

func (pa *base) NeutralPrice() float64 {
	switch pa.state {
	case uptrend:
		return pa.last_low

	case downtrend:
		return pa.last_high

	default:
		return 0
	}
}
