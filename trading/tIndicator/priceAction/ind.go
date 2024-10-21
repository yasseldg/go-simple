package priceAction

// private types

type ind struct {
	ts    int64
	state state
}

func newInd(ts int64, state state) ind {
	return ind{ts: ts, state: state}
}

func (i *ind) Ts() int64 {
	return i.ts
}

func (i *ind) State() state {
	return i.state
}

func (i *ind) Bshf() []float64 {
	switch i.state {
	case uptrend:
		return []float64{1, 0}

	case downtrend:
		return []float64{0, 1}

	case neutral_up, neutral_down:
		return []float64{1, 1}

	default:
		return []float64{0, 0}
	}
}

func (i *ind) Trend() []float64 {
	switch i.state {
	case downtrend:
		return []float64{0, 1, 0}

	case neutral_up:
		return []float64{0, 1, 4}

	case uptrend:
		return []float64{1, 0, 7}

	case neutral_down:
		return []float64{1, 0, 3}

	default:
		return []float64{0, 0, 0}
	}
}
