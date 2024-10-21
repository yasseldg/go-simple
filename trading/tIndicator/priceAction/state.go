package priceAction

type state int

const (
	downtrend state = iota + 1
	neutral_up
	uptrend
	neutral_down
)

func (s state) String() string {
	switch s {
	case downtrend:
		return "downtrend"
	case neutral_up:
		return "neutral_up"
	case uptrend:
		return "uptrend"
	case neutral_down:
		return "neutral_down"
	default:
		return ""
	}
}

func (s state) IsDowntrend() bool {
	return s == downtrend
}

func (s state) IsNeutralUp() bool {
	return s == neutral_up
}

func (s state) IsUptrend() bool {
	return s == uptrend
}

func (s state) IsNeutralDown() bool {
	return s == neutral_down
}
